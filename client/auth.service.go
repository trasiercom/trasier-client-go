package client

import (
	"net/http"
	"log"
	"io/ioutil"
	"net/url"
	"strings"
)

type AuthService struct {
	clientId     string
	clientSecret string
}

func NewAuthService(clientId string, clientSecret string) *AuthService {
	return &AuthService{clientId: clientId, clientSecret: clientSecret}
}

func (authService *AuthService) GetAuthToken() (authToken string) {
	client := &http.Client{}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Add("scope", "")

	req, err := http.NewRequest("POST", "https://trasier-keycloak-test.app.trasier.com/auth/realms/trasier-dev/protocol/openid-connect/token", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(authService.clientId, authService.clientSecret)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	return string(bodyText)
}
