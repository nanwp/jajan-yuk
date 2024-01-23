package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nanwp/jajan-yuk/payment/core/entity"
	"github.com/nanwp/jajan-yuk/payment/core/service"
)

type HttpHandler interface {
	GenerateSnapURL(w http.ResponseWriter, r *http.Request)
	VerifyPayment(w http.ResponseWriter, r *http.Request)
	InitializePayment(w http.ResponseWriter, r *http.Request)
	PaymentNotification(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	paymentService  service.PaymentService
	midtransService service.MidtransService
}

func NewHttpHandler(paymentService service.PaymentService, midtransService service.MidtransService) HttpHandler {
	return &httpHandler{
		paymentService:  paymentService,
		midtransService: midtransService,
	}
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *httpHandler) InitializePayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response
	// ctx := r.Context()

	if r.Method != http.MethodPost {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			response.Message = err.Error()
			response.Success = false
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		defer r.Body.Close()
	}

	params := entity.PaymentRequest{}
	err = json.Unmarshal(bodyBytes, &params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	payment, err := h.paymentService.InitializePayment(params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = payment
	response.Success = true
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func (h *httpHandler) GenerateSnapURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response
	// ctx := r.Context()

	if r.Method != http.MethodPost {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			response.Message = err.Error()
			response.Success = false
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		defer r.Body.Close()
	}

	params := entity.PaymentRequest{}
	err = json.Unmarshal(bodyBytes, &params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	payment, err := h.paymentService.InitializePayment(params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = payment
	response.Success = true
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
}

func (h *httpHandler) PaymentNotification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}
	var err error
	var response response

	if r.Method != http.MethodPost {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Body != nil {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.Message = err.Error()
			response.Success = false
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		defer r.Body.Close()

		err = json.Unmarshal(bodyBytes, &notificationPayload)
		if err != nil {
			response.Message = err.Error()
			response.Success = false
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	orderID, ok := notificationPayload["order_id"].(string)
	if !ok {
		response.Message = "order_id not found"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	success, err := h.midtransService.VerifyPayment(notificationPayload)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if success {
		_ = h.paymentService.ConfirmedPayment(orderID)
	}

	response.Success = true
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
