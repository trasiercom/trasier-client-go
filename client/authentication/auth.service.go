package authentication

import (
	"time"
	"sync"
)

const keycloakUrl = "https://trasier-keycloak-test.app.trasier.com/auth/realms/trasier-dev/protocol/openid-connect/token"

type AuthService struct {
	clientId            string
	clientSecret        string
	tokenService        *TokenService
	authToken           *AuthToken
	tokenExpirationTime time.Time
	mux                 sync.Mutex
}

func NewAuthService(clientId string, clientSecret string) *AuthService {
	tokenService := NewTokenService(clientId, clientSecret)
	return &AuthService{clientId: clientId, clientSecret: clientSecret, tokenService: tokenService}
}

func (authService *AuthService) GetToken() (authToken *AuthToken) {
	authService.mux.Lock()

	if authService.authToken == nil {
		authToken = authService.getNewAuthToken()
		authService.mux.Unlock()
		return
	}

	if authService.isTokenExpired() {
		// TODO refresh token
	}

	authToken = authService.authToken
	authService.mux.Unlock()
	return
}

func (authService *AuthService) getNewAuthToken() (authToken *AuthToken) {
	authToken = authService.tokenService.GetNewToken()
	authService.authToken = authToken
	timeToAdd := time.Second * time.Duration(authService.authToken.ExpiresIn)
	authService.tokenExpirationTime = time.Now().Add(timeToAdd)
	return
}

func (authService *AuthService) isTokenExpired() bool {
	return time.Now().After(authService.tokenExpirationTime)
}
