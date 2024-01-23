package repository

import "github.com/nanwp/jajan-yuk/payment/core/entity"

type PaymentRepository interface {
	GetPaymentByID(id string) (entity.Payment, error)
	CreatePayment(payment entity.Payment) (entity.Payment, error)
	UpdatePayment(payment entity.Payment) (entity.Payment, error)
}
