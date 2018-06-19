package client

import "../api"

type TrasierClient struct {
	spanService SpanService
	baseUrl     string
}

func NewTrasierClient(accountId string, spaceKey string) *TrasierClient {
	trasierClient := TrasierClient{}
	spanService := SpanService{AccountId: accountId, SpaceKey: spaceKey}
	trasierClient.spanService = spanService
	trasierClient.baseUrl = "http://localhost:8081/api/"
	return &trasierClient
}

func (t *TrasierClient) SendSpans(spans []api.Span) error {
	return t.spanService.SendSpans(spans, t.baseUrl)
}
