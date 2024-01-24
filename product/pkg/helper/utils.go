package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeResponseBody(resp *http.Response, v interface{}) error {
	// create decoder
	decoder := json.NewDecoder(resp.Body)

	// decode response body
	err := decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

// create func to convert http req to curls
func ReqToCurl(req *http.Request) (string, error) {
	// create curl
	curl := "curl -X " + req.Method

	// add headers
	for k, v := range req.Header {
		curl += " -H '" + k + ": " + v[0] + "'"
	}

	// add body
	if req.Body != nil {
		curl += fmt.Sprintf(" -d '%v'", req.Body)
	}

	// add url
	curl += " " + req.URL.String()

	return curl, nil
}
