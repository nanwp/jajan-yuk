package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nanwp/jajan-yuk/payment/core/entity"
	"github.com/nanwp/jajan-yuk/payment/core/repository"
)

type PaymentService interface {
	ConfirmedPayment(id string) error
	InitializePayment(params entity.PaymentRequest) (response entity.PaymentResponse, err error)
}

type paymentService struct {
	paymentRepo     repository.PaymentRepository
	midtransService MidtransService
}

func NewPaymentService(paymentRepo repository.PaymentRepository, midtransService MidtransService) PaymentService {
	return &paymentService{
		paymentRepo:     paymentRepo,
		midtransService: midtransService,
	}
}

func (s *paymentService) ConfirmedPayment(id string) error {
	payment, err := s.paymentRepo.GetPaymentByID(id)
	if err != nil {
		return err
	}

	if payment == (entity.Payment{}) {
		return errors.New("payment requestnot found")
	}

	payment.Status = "confirmed"
	_, err = s.paymentRepo.UpdatePayment(payment)
	if err != nil {
		return err
	}

	return nil
}

func (s *paymentService) InitializePayment(params entity.PaymentRequest) (response entity.PaymentResponse, err error) {
	payment := entity.Payment{
		ID:      uuid.New().String(),
		UserID:  params.UserID,
		Ammount: params.Ammount,
	}

	if err := s.midtransService.GenerateSnapURL(&payment); err != nil {
		return response, err
	}

	payment, err = s.paymentRepo.CreatePayment(payment)
	if err != nil {
		return response, err
	}

	return entity.PaymentResponse{
		SnapURL: payment.SnapURL,
	}, nil
}
