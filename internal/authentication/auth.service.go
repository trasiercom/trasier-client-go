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
		authToken = authService.requestNewToken()
		authService.mux.Unlock()
		return
	}

	if authService.isTokenExpired() {
		authToken = authService.refreshToken()
		authService.mux.Unlock()
		return
	}

	authToken = authService.authToken
	authService.mux.Unlock()
	return
}

func (authService *AuthService) requestNewToken() (authToken *AuthToken) {
	authToken = authService.tokenService.GetNewToken()
	authService.updateToken(authToken)
	return
}

func (authService *AuthService) refreshToken() (authToken *AuthToken) {
	authToken = authService.tokenService.RefreshToken(authService.authToken.RefreshToken)
	authService.updateToken(authToken)
	return
}

func (authService *AuthService) updateToken(authToken *AuthToken) {
	authService.authToken = authToken
	timeToAdd := time.Second * time.Duration(authService.authToken.ExpiresIn)
	authService.tokenExpirationTime = time.Now().Add(timeToAdd)
}

func (authService *AuthService) isTokenExpired() bool {
	return time.Now().After(authService.tokenExpirationTime)
}
