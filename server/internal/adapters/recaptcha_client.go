package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RecaptchaVerifyRequestBody struct {
	Secret   string `json:"secret"`
	Response string `json:"response"`
}

type RecaptchaVerifyResponse struct {
	Success    bool     `json:"success"`
	Timestamp  string   `json:"challenge_ts"`
	Hostname   string   `json:"hostname"`
	ErrorCodes []string `json:"error_codes"`
}

type RecaptchaClient interface {
	VerifyToken(ctx context.Context, token string) (bool, error)
}

type recaptchaClient struct {
	recaptchaSecret    string
	recaptchaVerifyURL string
}

func NewRecaptchaClient(recaptchaSecret string) RecaptchaClient {
	return &recaptchaClient{
		recaptchaSecret:    recaptchaSecret,
		recaptchaVerifyURL: "https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s",
	}
}

func (rc *recaptchaClient) VerifyToken(ctx context.Context, token string) (bool, error) {
	requestUrl := fmt.Sprintf(rc.recaptchaVerifyURL, rc.recaptchaSecret, token)

	res, err := http.Post(requestUrl, "application/json", nil)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	if len(body) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, body, "", "\t"); err != nil {
			fmt.Printf("JSON parse error: %v", err)
			return false, err
		}
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("Body: No Body Supplied\n")
	}

	verifyResponse := &RecaptchaVerifyResponse{}
	err = json.Unmarshal(body, verifyResponse)
	if err != nil {
		return false, err
	}

	return verifyResponse.Success, nil
}
