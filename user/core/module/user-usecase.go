package module

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	"github.com/nanwp/jajan-yuk/user/core/publisher"
	"github.com/nanwp/jajan-yuk/user/core/repository"
	"github.com/nanwp/jajan-yuk/user/pkg/helper"
)

type UserUsecase interface {
	Register(ctx context.Context, user entity.User) (response entity.User, err error)
	ActivateAccount(ctx context.Context, params entity.ActivateAccount) (response entity.User, err error)
	RequestResetPassword(ctx context.Context, params entity.RequestResetPassword) error
	ResetPassword(params entity.ResetPassword) (response entity.User, err error)
}

type userUsecase struct {
	cfg            config.Config
	userRepo       repository.UserRepository
	emailPublisher publisher.EmailPublisher
}

func (u userUsecase) ResetPassword(params entity.ResetPassword) (response entity.User, err error) {
	if err := params.Validate(); err != nil {
		return response, err
	}

	id, err := u.userRepo.GetTokenFromRedis(params.Token)
	if err != nil {
		return response, err
	}

	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return response, err
	}

	password, _ := helper.GeneratePasswordString(params.NewPassword)

	user.Password = password

	if err := u.userRepo.UpdateUser(user); err != nil {
		return response, err
	}

	if err := u.userRepo.DeleteTokenFromRedis(params.Token); err != nil {
		return response, err
	}

	response = user

	response.Password = ""
	return
}

func (u userUsecase) RequestResetPassword(ctx context.Context, params entity.RequestResetPassword) error {
	user, err := u.userRepo.GetUserByEmail(params.Email)
	if err != nil {
		return err
	}

	token := helper.RandomSerialString(32)

	if err := u.userRepo.StoredTokenToRedis(token, user.ID, 10); err != nil {
		return err
	}

	timeNow := helper.GetCurrentTime().Format("02/January/2006 15:04:05")
	timeExpired := helper.GetCurrentTime().Add(time.Minute * 10).Format("02/January/2006 15:04:05")

	email := entity.Email{
		Title:    fmt.Sprintf("Request Reset Password"),
		Receiver: user.Email,
		Subject:  fmt.Sprintf("Request Reset Password"),
		Body:     helper.ResetPassword(user.Name, fmt.Sprintf("%s/reset-password?token=%s", u.cfg.WebURL, token), timeNow, timeExpired),
	}

	err = u.emailPublisher.SendEmail(ctx, email)
	if err != nil {
		return err
	}

	return nil
}

func (u userUsecase) ActivateAccount(ctx context.Context, params entity.ActivateAccount) (response entity.User, err error) {
	if params.Token == "" {
		return response, fmt.Errorf("token required")
	}

	id, err := u.userRepo.GetTokenFromRedis(params.Token)
	if err != nil {
		return response, fmt.Errorf("invalid or expired token")
	}

	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return response, err
	}

	if user.ActivatedAt.Valid {
		return response, fmt.Errorf("user are already activated")
	}

	user.ActivatedAt.Valid = true
	user.ActivatedAt.Time = time.Now()

	err = u.userRepo.UpdateUser(user)
	if err != nil {
		return response, err
	}

	email := entity.Email{
		Title:    fmt.Sprintf("Successful activate account"),
		Receiver: user.Email,
		Subject:  fmt.Sprintf("Activate account"),
		Body:     helper.SuccesActivateEmail(user.Name),
	}

	if err := u.emailPublisher.SendEmail(ctx, email); err != nil {
		log.Printf("error at %v", err)
	}

	response = user
	return
}

func (u userUsecase) Register(ctx context.Context, user entity.User) (response entity.User, err error) {
	if err := user.Validate(); err != nil {
		return response, err
	}

	role, err := u.userRepo.GetRoleByID(user.Role.ID)
	if err != nil {
		return response, err
	}

	user.Role = role

	usernameUsed, err := u.userRepo.CheckUsername(user.Username)
	if err != nil {
		return response, err
	}

	if usernameUsed {
		return response, entity.ErrorUsernameUsed
	}

	emailUsed, err := u.userRepo.CheckEmail(user.Email)
	if err != nil {
		return response, err
	}

	if emailUsed {
		return response, entity.ErrorEmailUsed
	}

	password, err := helper.GeneratePasswordString(user.Password)
	if err != nil {
		errMsg := fmt.Errorf("error when generate password: %v", err.Error())
		return response, errMsg
	}

	user.Password = password

	response, err = u.userRepo.AddUser(user)
	if err != nil {
		return response, err
	}

	token := helper.RandomSerialString(32)
	if err := u.userRepo.StoredTokenToRedis(token, response.ID, 30); err != nil {
		return response, err
	}

	url := fmt.Sprintf("%s/verification?token=%s", u.cfg.WebURL, token)
	body := helper.RegisterEmail(user.Name, strings.ToLower(user.Role.Name), url)

	email := entity.Email{
		Title:    fmt.Sprintf("Registration %v - Jajan Yuk Apps", strings.ToLower(user.Role.Name)),
		Receiver: user.Email,
		Subject:  "Jajan Yuk - User Verification",
		Body:     body,
	}

	err = u.emailPublisher.SendEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}

	response.Role = role
	response.Password = ""

	return response, nil
}

func NewUserRepository(cfg config.Config, userRepo repository.UserRepository, emailPublisher publisher.EmailPublisher) UserUsecase {
	return userUsecase{
		cfg:            cfg,
		userRepo:       userRepo,
		emailPublisher: emailPublisher,
	}
}
