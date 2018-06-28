package client

import "../api"

type TrasierClient struct {
	spanService SpanService
	baseUrl     string
}

func NewTrasierClient() *TrasierClient {
	trasierClient := TrasierClient{}
	spanService := SpanService{}
	trasierClient.spanService = spanService
	trasierClient.baseUrl = "http://localhost:8081/api/"
	return &trasierClient
}

func (t *TrasierClient) SendSpans(accountId string, spaceKey string, spans []api.Span) error {
	return t.spanService.SendSpans(accountId, spaceKey, spans, t.baseUrl)
}
