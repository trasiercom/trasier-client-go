package authentication

import (
	"log"
	"encoding/json"
	"time"
)

const keycloakUrl = "https://trasier-keycloak-test.app.trasier.com/auth/realms/trasier-dev/protocol/openid-connect/token"

type AuthService struct {
	clientId            string
	clientSecret        string
	tokenService        *TokenService
	authToken           *AuthToken
	tokenExpirationTime time.Time
}

func NewAuthService(clientId string, clientSecret string) *AuthService {
	tokenService := NewTokenService(clientId, clientSecret)
	return &AuthService{clientId: clientId, clientSecret: clientSecret, tokenService: tokenService}
}

func (authService *AuthService) GetToken() *AuthToken {

	// TODO lock with mutex

	if authService.authToken == nil {
		return authService.getNewAuthToken()
	}

	if (authService.isTokenExpired()) {
		// TODO refresh token
	}

	return authService.authToken
}

func (authService *AuthService) getNewAuthToken() *AuthToken {
	body := authService.tokenService.requestToken()
	authService.authToken = authService.convertToAuthToken(body)
	timeToAdd := time.Second * time.Duration(authService.authToken.ExpiresIn)
	authService.tokenExpirationTime = time.Now().Add(timeToAdd)
	return authService.authToken
}

func (authService *AuthService) convertToAuthToken(body []byte) *AuthToken {
	authToken := new(AuthToken)
	err := json.Unmarshal(body, &authToken)
	if err != nil {
		log.Fatal("Something went wrong during unmarshalling the response", err.Error())
	}
	return authToken
}

func (authService *AuthService) isTokenExpired() bool {
	return time.Now().After(authService.tokenExpirationTime)
}
