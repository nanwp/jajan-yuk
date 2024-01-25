package product_client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nanwp/jajan-yuk/pedagang/config"
	"github.com/nanwp/jajan-yuk/pedagang/pkg/helper"
)

type ProductClient interface {
	GetProductByUserID(userID string) (GetPedagangByUserIDResponse, error)
}

type productClient struct {
	cfg config.Config
}

func NewProductClient(cfg config.Config) ProductClient {
	return &productClient{
		cfg: cfg,
	}
}

func (c *productClient) GetProductByUserID(userID string) (GetPedagangByUserIDResponse, error) {
	url := fmt.Sprintf("%s/api/v1/product/pedagang", c.cfg.ProductURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errMsg := fmt.Errorf("[ProductClient.GetPedagangByUserID] error when create new request: %w", err)
		return GetPedagangByUserIDResponse{}, errMsg
	}

	req.Header.Set("Secret-Key", c.cfg.SecretKey)

	q := req.URL.Query()
	q.Add("user_id", userID)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errMsg := fmt.Errorf("[ProductClient.GetPedagangByUserID] error when send request: %w", err)
		return GetPedagangByUserIDResponse{}, errMsg
	}

	defer resp.Body.Close()

	log.Println("resp status code:", resp.StatusCode)
	curl, _ := helper.ReqToCurl(req)
	log.Printf("resp status: %v", curl)

	var result GetPedagangByUserIDResponse
	err = helper.DecodeResponseBody(resp, &result)
	if err != nil {
		log.Println("error when decode response body:", err)
		errMsg := fmt.Errorf("[ProductClient.GetPedagangByUserID] error when read response body: %w", err)
		return GetPedagangByUserIDResponse{}, errMsg
	}

	return result, err
}
