package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nanwp/jajan-yuk/pedagang/core/entity"
	"github.com/nanwp/jajan-yuk/pedagang/core/module"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/helper"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/lib/auth"
)

type HTTPHandler interface {
	GetPedagangByID(w http.ResponseWriter, r *http.Request)
	GetAllPedagangNearby(w http.ResponseWriter, r *http.Request)
	CreatePedagang(w http.ResponseWriter, r *http.Request)
	UpdateLocation(w http.ResponseWriter, r *http.Request)
	SwitchStatus(w http.ResponseWriter, r *http.Request)
	GetImage(w http.ResponseWriter, r *http.Request)
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

type httpHandler struct {
	pedagangService module.PedagangService
}

func NewHTTPHandler(pedagangService module.PedagangService) HTTPHandler {
	return &httpHandler{
		pedagangService: pedagangService,
	}
}

func (h *httpHandler) GetImage(w http.ResponseWriter, r *http.Request) {
	imageName := mux.Vars(r)["image_name"]

	_, err := os.Open("./images/" + imageName)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, "./images/"+imageName)
	return
}

func (h *httpHandler) SwitchStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var err error
	var response response
	var bodyBytes []byte

	//auth
	user, err := auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		bodyBytes, err = response.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	//get body

	//validate

	//switch status
	err = h.pedagangService.SwitchActiveStatus(user.Data.User.ID)
	if err != nil {
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

	response.Success = true
	response.Message = "Success switch status"
	bodyBytes, err = response.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) GetPedagangByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	//auth
	_, err = auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		bodyBytes, err = response.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	//get params url /:id
	id := r.URL.Query().Get("id")

	//validate

	//get pedagang by id
	pedagang, err := h.pedagangService.GetPedagangByID(id)
	if err != nil {
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

	response.Success = true
	response.Message = "Success get pedagang by id"
	response.Data = pedagang
	bodyBytes, err = response.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) GetAllPedagangNearby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response

	//auth
	_, err = auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		bodyBytes, err = response.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	//get body
	var params entity.GetAllPedagangNearbyRequest
	params.Keyword = r.URL.Query().Get("keyword")
	params.Latitude = helper.StringToFloat64(r.URL.Query().Get("latitude"))
	params.Longitude = helper.StringToFloat64(r.URL.Query().Get("longitude"))
	params.MaxDistance = helper.StringToFloat64(r.URL.Query().Get("max_distance"))

	//validate

	//get all pedagang nearby
	pedagangs, err := h.pedagangService.GetAllPedagangNearby(params)
	if err != nil {
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

	response.Success = true
	response.Message = "Success get all pedagang nearby"
	response.Data = pedagangs
	bodyBytes, err = response.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return

}

func (h *httpHandler) CreatePedagang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response
	log.Println("masuk create pedagang")
	log.Println(helper.ReqToCurl(r))

	//auth
	_, err = auth.ValidateCurrentUser(w, r)
	if err != nil {
		log.Printf("[httpHandler.CreatePedagang] error when validate current user: %v\n", err)
		response.Message = err.Error()
		bodyBytes, err = response.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	//get body
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Printf("[httpHandler.CreatePedagang] error when parse multipart form: %v\n", err)
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

	var params entity.Pedagang
	params.UserID = r.FormValue("user_id")
	params.NameMerchant = r.FormValue("name_merchant")
	params.Phone = r.FormValue("phone")
	log.Println(params)

	//upload image
	//
	file, handler, err := r.FormFile("image")
	log.Println(file)
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

		imagePath, err := h.pedagangService.SaveImage(file, handler)
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
		params.Image = imagePath
	}

	//validate

	//create pedagang
	pedagang, err := h.pedagangService.CreatePedagang(params)
	if err != nil {
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

	response.Success = true
	response.Message = "Success create pedagang"
	response.Data = pedagang
	bodyBytes, err = response.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}

func (h *httpHandler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bodyBytes []byte
	var err error
	var response response
	// ctx := r.Context()

	//auth
	user, err := auth.ValidateCurrentUser(w, r)
	if err != nil {
		response.Message = err.Error()
		bodyBytes, err = response.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(bodyBytes)
		return
	}

	//get body
	var params entity.UpdateLocationRequest
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
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

	params.UserID = user.Data.User.ID

	//update location
	err = h.pedagangService.UpdateLocation(params)
	if err != nil {
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

	response.Success = true
	response.Message = "Success update location"
	bodyBytes, err = response.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
	return
}
