package authentication

import (
	"log"
	"net/http"
	"strings"
	"net/url"
	"io/ioutil"
)

type TokenService struct {
	request *http.Request
}

func NewTokenService(clientId string, clientSecret string) *TokenService {
	request, err := createRequest(clientId, clientSecret)
	if err != nil {
		log.Fatal("Something went wrong while creating the request", err)
	}
	return &TokenService{request: request}
}

func (tokenService *TokenService) requestToken() []byte {
	client := &http.Client{}
	resp, err := client.Do(tokenService.request)
	if err != nil {
		log.Fatal("Something went wrong while sending the request", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Something went wrong whily reading the response", err.Error())
	}
	return body
}

func createRequest(clientId string, clientSecret string) (request *http.Request, err error) {
	request, err = http.NewRequest("POST", keycloakUrl, strings.NewReader(createRequestBody().Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(clientId, clientSecret)
	return
}

func createRequestBody() (requestBody url.Values) {
	requestBody = url.Values{}
	requestBody.Set("grant_type", "client_credentials")
	requestBody.Add("scope", "")
	return
}
