package pkg

import (
	"github.com/trasiercom/trasier-client-go/internal/spans"
	"github.com/trasiercom/trasier-client-go/api"
)

type TrasierClient struct {
	spanService spans.SpanService
	baseUrl     string
}

func NewTrasierClient() *TrasierClient {
	trasierClient := TrasierClient{}
	trasierClient.spanService = spans.SpanService{}
	trasierClient.baseUrl = "http://localhost:8081/api/"
	return &trasierClient
}

func (t *TrasierClient) SendSpans(accountId string, spaceKey string, spans []api.Span) error {
	return t.spanService.SendSpans(accountId, spaceKey, spans, t.baseUrl)
}
