package pedagang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nanwp/jajan-yuk/user/config"
	"github.com/nanwp/jajan-yuk/user/core/entity"
	"github.com/nanwp/jajan-yuk/user/pkg/helper"
)

type PedagangClient interface {
	CreatePedagang(pedagang entity.Pedagang) (entity.Pedagang, error)
}

type pedagangClient struct {
	cfg config.Config
}

func NewPedagangClient(cfg config.Config) PedagangClient {
	return &pedagangClient{
		cfg: cfg,
	}
}

func (p *pedagangClient) CreatePedagang(pedagang entity.Pedagang) (entity.Pedagang, error) {
	url := fmt.Sprintf("%s/api/v1/pedagang", p.cfg.PedagangURL)

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	writer.WriteField("user_id", pedagang.UserID)
	writer.WriteField("name_merchant", pedagang.NameMerchant)
	writer.WriteField("phone", pedagang.Phone)

	// Open the existing file
	file, err := os.Open(pedagang.Image)
	if err != nil {
		return entity.Pedagang{}, err
	}
	defer file.Close()

	// Create a form file writer for the image file
	fileWriter, err := writer.CreateFormFile("image", filepath.Base(pedagang.Image))
	if err != nil {
		return entity.Pedagang{}, err
	}

	// Copy the file data to the writer
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return entity.Pedagang{}, err
	}

	err = writer.Close()
	if err != nil {
		return entity.Pedagang{}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return entity.Pedagang{}, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Secret-Key", p.cfg.SecretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return entity.Pedagang{}, err
	}

	defer resp.Body.Close()

	var response CreatePedagangResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return entity.Pedagang{}, err
	}

	if err := helper.DeleteImage(pedagang.Image); err != nil {
		log.Println(err)
	}

	if response.Success == false {
		return entity.Pedagang{}, fmt.Errorf(response.Message)
	}

	return response.Data.ToEntity(), nil
}
