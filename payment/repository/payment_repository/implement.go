package payment_repository

import (
	"github.com/nanwp/jajan-yuk/payment/config"
	"github.com/nanwp/jajan-yuk/payment/core/entity"
	repository_intf "github.com/nanwp/jajan-yuk/payment/core/repository"
	"gorm.io/gorm"
)

type repository struct {
	cfg config.Config
	db  *gorm.DB
}

func New(cfg config.Config, db *gorm.DB) repository_intf.PaymentRepository {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *repository) GetPaymentByID(id string) (entity.Payment, error) {
	db := r.db.Model(Payment{}).Where("id = ?", id)

	var payment Payment
	if err := db.First(&payment).Error; err != nil {
		return entity.Payment{}, err
	}

	return payment.ToEntoty(), nil
}

func (r *repository) CreatePayment(payment entity.Payment) (entity.Payment, error) {
	db := r.db.Model(Payment{})
	paymentModel := Payment{}
	paymentModel.FromEntoty(payment)
	paymentModel.CreatedBy = payment.UserID
	paymentModel.UpdatedBy = payment.UserID

	result := db.Create(&paymentModel)
	if result.Error != nil {
		return entity.Payment{}, result.Error
	}

	return paymentModel.ToEntoty(), nil
}

func (r *repository) UpdatePayment(payment entity.Payment) (entity.Payment, error) {
	db := r.db.Model(Payment{}).Where("id = ?", payment.ID)
	paymentModel := Payment{}
	paymentModel.FromEntoty(payment)
	paymentModel.UpdatedBy = payment.UserID

	result := db.Updates(&paymentModel)
	if result.Error != nil {
		return entity.Payment{}, result.Error
	}

	return paymentModel.ToEntoty(), nil
}
