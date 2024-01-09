package module

import (
	"context"
	"fmt"
	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	"github.com/nanwp/jajan-yuk/user/core/publisher"
	"github.com/nanwp/jajan-yuk/user/core/repository"
	"github.com/nanwp/jajan-yuk/user/pkg/helper"
	"strings"
)

type UserUsecase interface {
	Register(ctx context.Context, user entity.User) (response entity.User, err error)
}

type userUsecase struct {
	cfg            config.Config
	userRepo       repository.UserRepository
	emailPublisher publisher.EmailPublisher
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

	url := fmt.Sprintf("%s/activate?token=%s", u.cfg.BaseUrl, token)
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

	return response, nil
}

func NewUserRepository(cfg config.Config, userRepo repository.UserRepository, emailPublisher publisher.EmailPublisher) UserUsecase {
	return userUsecase{
		cfg:            cfg,
		userRepo:       userRepo,
		emailPublisher: emailPublisher,
	}
}
