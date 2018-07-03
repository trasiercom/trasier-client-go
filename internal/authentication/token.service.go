package authentication

import (
	"log"
	"net/http"
	"strings"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type TokenService struct {
	clientId     string
	clientSecret string
}

func NewTokenService(clientId string, clientSecret string) *TokenService {
	return &TokenService{clientId: clientId, clientSecret: clientSecret}
}

func (tokenService *TokenService) GetNewToken() *AuthToken {
	requestBody := url.Values{}
	requestBody.Set("grant_type", "client_credentials")
	requestBody.Add("scope", "")
	tokenValue := tokenService.requestToken(requestBody)
	return tokenService.convertToAuthToken(tokenValue)
}

func (tokenService *TokenService) RefreshToken(refreshToken string) *AuthToken {
	requestBody := url.Values{}
	requestBody.Set("grant_type", "refresh_token")
	requestBody.Add("refresh_token", refreshToken)
	tokenValue := tokenService.requestToken(requestBody)
	return tokenService.convertToAuthToken(tokenValue)
}

func (tokenService *TokenService) requestToken(requestBody url.Values) []byte {
	request, err := createBasicAuthRequest(tokenService.clientId, tokenService.clientSecret, requestBody)
	if err != nil {
		log.Fatal("Something went wrong while creating the request", err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal("Something went wrong while sending the request", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Something went wrong whily reading the response", err.Error())
	}
	return body
}

func createBasicAuthRequest(clientId string, clientSecret string, requestBody url.Values) (request *http.Request, err error) {
	request, err = http.NewRequest("POST", keycloakUrl, strings.NewReader(requestBody.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(clientId, clientSecret)
	return
}

func (tokenService *TokenService) convertToAuthToken(body []byte) *AuthToken {
	authToken := new(AuthToken)
	err := json.Unmarshal(body, &authToken)
	if err != nil {
		log.Fatal("Something went wrong during unmarshalling the response", err.Error())
	}
	return authToken
}
