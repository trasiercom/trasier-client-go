package impl

import (
	"net/url"
	"net/http"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	httpClient *http.Client
	SpanService *SpanService
}

func TrasierClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client {httpClient: httpClient}
	c.SpanService = &SpanService{client: c}
}