package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nanwp/jajan-yuk/pedagang/config"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/helper"
)

var (
	cfg = config.Get()
)

func validateSecretKey(secretKey string) (record SecretKey, err error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/validate-secret-key", cfg.AUTH_URL), nil)
	if err != nil {
		return
	}
	req.Header.Set("Secret-Key", secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	bodyString := string(body)

	httpResp := ValidateSecretKeyResponse{}
	err = json.Unmarshal([]byte(bodyString), &httpResp)
	if err != nil {
		return
	}

	if httpResp.Success == false {
		return record, fmt.Errorf(httpResp.Message)
	}

	return httpResp.Data, nil
}

func ValidateCurrentUser(w http.ResponseWriter, r *http.Request) (user GetCurrentUserResponse, err error) {

	secretKey := r.Header.Get("Secret-Key")

	if secretKey != "" {
		record, err := validateSecretKey(secretKey)
		if err != nil {
			return user, err
		}

		if record.ID == "" {
			return user, fmt.Errorf("Secret key not found")
		}

		simulatedUser := User{
			ID:       record.ID,
			Username: record.Serial,
			Email:    record.Name,
			Role:     record.Role,
		}

		user = GetCurrentUserResponse{
			Success: true,
			Message: "Success get current user",
			Data: CurrentUserData{
				User: simulatedUser,
			},
		}

		return user, nil
	}

	token := r.Header.Get("Authorization")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/current-user", cfg.AUTH_URL), nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", token)

	curl, _ := helper.ReqToCurl(req)
	fmt.Println(curl)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	bodyString := string(body)

	err = json.Unmarshal([]byte(bodyString), &user)
	if err != nil {
		return
	}

	log.Println(user.Data)

	if user.Data.User.ID == "" {
		return user, fmt.Errorf("User not found")
	}

	if user.Success == false {
		return user, fmt.Errorf(user.Message)
	}

	return user, nil
}
