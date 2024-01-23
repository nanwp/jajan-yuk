package service

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/nanwp/jajan-yuk/payment/config"
	"github.com/nanwp/jajan-yuk/payment/core/entity"
)

type MidtransService interface {
	GenerateSnapURL(payment *entity.Payment) error
	VerifyPayment(data map[string]interface{}) (bool, error)
}

type midtransService struct {
	cfg    config.Config
	client snap.Client
}

func NewMidtransService(cfg config.Config) MidtransService {
	var client snap.Client
	envi := midtrans.Sandbox
	if cfg.IsProduction {
		envi = midtrans.Production
	}

	client.New(cfg.MidtransKey, envi)

	return &midtransService{
		cfg:    cfg,
		client: client,
	}
}

func (s *midtransService) GenerateSnapURL(payment *entity.Payment) error {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.ID,
			GrossAmt: payment.Ammount,
		},
	}

	snapResp, err := s.client.CreateTransaction(req)
	if err != nil {
		errMsg := fmt.Errorf("[MidtransService.GenerateSnapURL] error when create transaction: %w", err)
		return errMsg
	}

	payment.SnapURL = snapResp.RedirectURL
	return nil
}

func (s *midtransService) VerifyPayment(data map[string]interface{}) (bool, error) {
	var client coreapi.Client
	envi := midtrans.Sandbox
	if s.cfg.IsProduction {
		envi = midtrans.Production
	}

	client.New(s.cfg.MidtransKey, envi)

	orderID, ok := data["order_id"].(string)
	if !ok {
		errMsg := fmt.Errorf("[MidtransService.VerifyPayment] error when get order_id")
		return false, errMsg
	}

	transactionStatusResp, err := client.CheckTransaction(orderID)
	if err != nil {
		errMsg := fmt.Errorf("[MidtransService.VerifyPayment] error when check transaction: %w", err)
		return false, errMsg
	}

	if transactionStatusResp != nil {
		if transactionStatusResp.TransactionStatus == "capture" {
			if transactionStatusResp.FraudStatus == "challenge" {
				fmt.Println("Transaction is challenged")
			} else if transactionStatusResp.FraudStatus == "accept" {
				return true, nil
			}
		} else if transactionStatusResp.TransactionStatus == "settlement" {
			return true, nil
		} else if transactionStatusResp.TransactionStatus == "deny" {
			fmt.Println("Transaction is denied")
		} else if transactionStatusResp.TransactionStatus == "cancel" {
			fmt.Println("Transaction is canceled")
		} else if transactionStatusResp.TransactionStatus == "expire" {
			fmt.Println("Transaction is expired")
		} else if transactionStatusResp.TransactionStatus == "pending" {
			fmt.Println("Transaction is pending")
		}
	}

	return false, nil
}
