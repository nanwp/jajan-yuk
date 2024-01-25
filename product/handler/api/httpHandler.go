package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nanwp/jajan-yuk/product/core/entity"
	"github.com/nanwp/jajan-yuk/product/core/module"
	"github.com/nanwp/jajan-yuk/product/pkg/helper"
	"github.com/nanwp/jajan-yuk/product/pkg/lib/auth"
)

type HTTPHandler interface {
	GetCategoryByIDs(w http.ResponseWriter, r *http.Request)
	GetCategoryByUserCreated(w http.ResponseWriter, r *http.Request)
	CreateCategory(w http.ResponseWriter, r *http.Request)
	GetVariantByIDs(w http.ResponseWriter, r *http.Request)
	GetVariantByUserCreated(w http.ResponseWriter, r *http.Request)
	CreateVariant(w http.ResponseWriter, r *http.Request)
}

type httpHandler struct {
	variantService  module.VariantService
	categoryService module.CategoryService
}

func NewHTTPHandler(variantService module.VariantService, categoryService module.CategoryService) HTTPHandler {
	return &httpHandler{
		variantService:  variantService,
		categoryService: categoryService,
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

func (h *httpHandler) GetCategoryByIDs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	_, err = auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	idString := r.URL.Query().Get("ids")

	ids := strings.Split(idString, ",")
	if len(ids) == 0 {
		response.Message = "ids is required"
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bodyBytes)
		return
	}

	idsInt, err := helper.ConvertStringSliceToInt64Slice(ids)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bodyBytes)
		return
	}

	records, err := h.categoryService.GetCategoryByIDs(idsInt)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bodyBytes)
		return
	}

	response.Message = "Success get category by ids"
	response.Success = true
	response.Data = records
	bodyBytes, err = json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) GetCategoryByUserCreated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	user, err := auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	records, err := h.categoryService.GetCategoryByUserCreated(user.Data.User.ID)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bodyBytes)
		return
	}

	response.Message = "Success get category by user created"
	response.Success = true
	response.Data = records
	bodyBytes, err = json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	user, err := auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	var category entity.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bodyBytes)
		return
	}

	record, err := h.categoryService.CreateCategory(category, user.Data.User.ID)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bodyBytes)
		return
	}

	response.Message = "Success create category"
	response.Success = true
	response.Data = record
	bodyBytes, err = json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) CreateVariant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	user, err := auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	var variant entity.Variant
	err = json.NewDecoder(r.Body).Decode(&variant)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bodyBytes)
		return
	}

	record, err := h.variantService.CreateVariant(variant, user.Data.User.ID)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bodyBytes)
		return
	}

	response.Message = "Success create variant"
	response.Success = true
	response.Data = record
	bodyBytes, err = json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) GetVariantByUserCreated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	user, err := auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	records, err := h.variantService.GetVariantByUserCreated(user.Data.User.ID)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bodyBytes)
		return
	}

	response.Message = "Success get variant by user created"
	response.Success = true
	response.Data = records
	bodyBytes, err = json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) GetVariantByIDs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	_, err = auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	idString := r.URL.Query().Get("ids")

	ids := strings.Split(idString, ",")

	if len(ids) == 0 {
		response.Message = "ids is required"
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bodyBytes)
		return
	}

	idsInt, err := helper.ConvertStringSliceToInt64Slice(ids)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(bodyBytes)
		return
	}

	records, err := h.variantService.GetVariantByIDs(idsInt)
	if err != nil {
		response.Message = err.Error()
		response.Success = false

		bodyBytes, err = json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bodyBytes)
		return
	}

	response.Message = "Success get variant by ids"
	response.Success = true
	response.Data = records
	bodyBytes, err = json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}
