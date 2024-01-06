package api

import (
	"encoding/json"
	"github.com/nanwp/jajan-yuk/auth/core/entity"
	"github.com/nanwp/jajan-yuk/auth/core/module"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetCurrentUser(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	authUsecase module.AuthUsecase
}

func NewHttpHandler(authUsecase module.AuthUsecase) HttpHandler {
	return &httpHandler{
		authUsecase: authUsecase,
	}
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h *httpHandler) Login(w http.ResponseWriter, r *http.Request) {
	var bodyBytes []byte
	var err error
	var response response
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		defer r.Body.Close()
	}

	loginRequest := entity.LoginRequest{}
	err = json.Unmarshal(bodyBytes, &loginRequest)
	if err != nil {
		log.Println(err)
		return
	}

	response.Data, err = h.authUsecase.Login(loginRequest)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Data = nil
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response.Success = true
		response.Message = "Login Succes"
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	var response response
	var err error
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		response.Success = false
		response.Message = "token required"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data, err = h.authUsecase.GetCurrentUser(token)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Data = nil
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response.Success = true
		response.Message = "success"
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var response response
	var err error
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		response.Success = false
		response.Message = "token required"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	params := entity.RefreshTokenRequest{RefreshToken: token}
	response.Data, err = h.authUsecase.RefreshToken(params)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response.Success = true
		response.Message = "success"
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(response)
}
