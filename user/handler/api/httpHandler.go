package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	"github.com/nanwp/jajan-yuk/user/core/module"
	"io/ioutil"
	"net/http"
)

type HttpHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	userUsecase module.UserUsecase
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (h httpHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response
	ctx := r.Context()

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

	user := entity.User{}
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	role := mux.Vars(r)["role"]

	if role == "USER" {
		user.Role.ID = "3e76048d-f9f2-4974-845f-60137f9e2f4b"
	} else if role == "PEDAGANG" {
		user.Role.ID = "ea8e1e87-ae6e-44b1-9854-4dbb0c70a330"
	} else {
		response.Message = "bad request"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data, err = h.userUsecase.Register(ctx, user)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		response.Data = nil
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response.Success = true
		response.Message = "register success, please check email"
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(response)
}

func NewHttpHandler(userUsecase module.UserUsecase) HttpHandler {
	return &httpHandler{
		userUsecase: userUsecase,
	}
}
