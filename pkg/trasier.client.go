package pkg

import (
	"github.com/trasiercom/trasier-client-go/internal/spans"
	"github.com/trasiercom/trasier-client-go/api"
	"github.com/trasiercom/trasier-client-go/internal/authentication"
	"os"
)

const TrasierClientIdKey = "trasier-client-id"
const TrasierClientSecretKey = "trasier-client-secret"

type trasierClient struct {
	spanService spans.SpanService
	baseUrl     string
	authService *authentication.AuthService
}

func NewTrasierClient() *trasierClient {
	trasierClient := trasierClient{}
	trasierClient.spanService = spans.SpanService{}
	trasierClient.baseUrl = "http://localhost:8081/api/"

	clientId := os.Getenv(TrasierClientIdKey)
	clientSecret := os.Getenv(TrasierClientSecretKey)
	trasierClient.authService = authentication.NewAuthService(clientId, clientSecret)

	return &trasierClient
}

func (t *trasierClient) SendSpans(accountId string, spaceKey string, spans []api.Span) {
	authToken := t.authService.GetToken()
	t.spanService.SendSpans(accountId, spaceKey, spans, t.baseUrl, authToken.AccessToken)
}
