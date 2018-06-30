package authentication

import (
	"net/http"
	"log"
	"net/url"
	"strings"
	"encoding/json"
	"io/ioutil"
)

const keycloakUrl = "https://trasier-keycloak-test.app.trasier.com/auth/realms/trasier-dev/protocol/openid-connect/token"

type AuthService struct {
	clientId     string
	clientSecret string
	request      *http.Request
}

func NewAuthService(clientId string, clientSecret string) *AuthService {
	request, err := createRequest(clientId, clientSecret)
	if err != nil {
		log.Fatal("Something went wrong while creating the request", err)
	}
	return &AuthService{clientId: clientId, clientSecret: clientSecret, request: request}
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

func (authService *AuthService) GetAuthToken() (*AuthToken) {
	body := authService.requestToken()
	return authService.convertToAuthToken(body)
}

func (authService *AuthService) convertToAuthToken(body []byte) *AuthToken {
	authToken := new(AuthToken)
	err := json.Unmarshal(body, &authToken)
	if err != nil {
		log.Fatal("Something went wrong during unmarshalling the response", err.Error())
	}
	return authToken
}

func (authService *AuthService) requestToken() []byte {
	client := &http.Client{}
	resp, err := client.Do(authService.request)
	if err != nil {
		log.Fatal("Something went wrong while sending the request", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Something went wrong whily reading the response", err.Error())
	}
	return body
}
