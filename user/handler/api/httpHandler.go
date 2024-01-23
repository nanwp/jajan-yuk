package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	"github.com/nanwp/jajan-yuk/user/core/module"
	"github.com/nanwp/jajan-yuk/user/pkg/helper"
)

type HttpHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Verification(w http.ResponseWriter, r *http.Request)
	RequestResetPassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	userUsecase module.UserUsecase
}

func NewHttpHandler(userUsecase module.UserUsecase) HttpHandler {
	return &httpHandler{
		userUsecase: userUsecase,
	}
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r *response) MarshalJSON() ([]byte, error) {
	type Alias response
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	})
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

	role := mux.Vars(r)["role"]

	if role == "user" {
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
		user.Role.ID = "3e76048d-f9f2-4974-845f-60137f9e2f4b"
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
	} else if role == "pedagang" {
		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			log.Println(err)
			response.Message = err.Error()
			bodyBytes, err = response.MarshalJSON()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(bodyBytes)
			return
		}

		user := entity.User{
			Name:     r.FormValue("name"),
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
			Address:  r.FormValue("address"),
			Image:    r.FormValue("image"),
			Role: entity.Role{
				ID: "ea8e1e87-ae6e-44b1-9854-4dbb0c70a330",
			},
			Pedagang: entity.Pedagang{
				NameMerchant: r.FormValue("name_merchant"),
				Phone:        r.FormValue("phone"),
			},
		}

		file, handler, err := r.FormFile("image")
		if err != nil {
			if err != http.ErrMissingFile {
				log.Println(err)
				response.Message = err.Error()
				bodyBytes, err = response.MarshalJSON()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write(bodyBytes)
				return
			}
		} else {
			defer file.Close()
			imagePath, err := helper.UploadImage(file, handler)
			if err != nil {
				log.Println(err)
				response.Message = err.Error()
				bodyBytes, err = response.MarshalJSON()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				w.Write(bodyBytes)
				return
			}
			user.Pedagang.Image = imagePath
		}
		response.Data, err = h.userUsecase.Register(ctx, user)
		if err != nil {
			response.Success = false
			response.Message = err.Error()
			response.Data = nil
			w.WriteHeader(http.StatusBadRequest)

			bodyBytes, err = response.MarshalJSON()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(bodyBytes)
			return
		} else {
			response.Success = true
			response.Message = "register success, please check email"
			w.WriteHeader(http.StatusOK)
		}

	} else {
		response.Message = "bad request"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h httpHandler) Verification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response
	ctx := r.Context()

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			if err != nil {
				response.Message = err.Error()
				response.Success = false
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}

			defer r.Body.Close()
		}
	}

	tkn := entity.ActivateAccount{}
	err = json.Unmarshal(bodyBytes, &tkn)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Method != http.MethodPost {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.userUsecase.ActivateAccount(ctx, tkn)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = user
	response.Success = true
	response.Message = "success activate account"
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func (h httpHandler) RequestResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response
	ctx := r.Context()

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			if err != nil {
				response.Message = err.Error()
				response.Success = false
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}

			defer r.Body.Close()
		}
	}

	params := entity.RequestResetPassword{}
	err = json.Unmarshal(bodyBytes, &params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Method != http.MethodPost {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.userUsecase.RequestResetPassword(ctx, params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Success = true
	response.Message = "success request reset password, please check email"
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func (h httpHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			if err != nil {
				response.Message = err.Error()
				response.Success = false
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}

			defer r.Body.Close()
		}
	}

	params := entity.ResetPassword{}
	err = json.Unmarshal(bodyBytes, &params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Method != http.MethodPost {
		response.Message = "Method not Allow"
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.userUsecase.ResetPassword(params)
	if err != nil {
		response.Message = err.Error()
		response.Success = false
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data = data
	response.Success = true
	response.Message = "success reset password, you can login with new password"
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}
