package impl

import (
	"net/url"
	"net/http"
	"encoding/json"
	"../api"
	"io"
	"bytes"
)

type SpanService struct {
	client *Client
}

func (c *SpanService) sendSpans(accountId string, spaceKey string, spans []api.Span) (bool, error) {
	req, err := c.newRequest("POST", "/users", nil)
	if err != nil {
		return false, err
	}
	var successful bool
	_, err = c.do(req, &successful)
	return successful, err
}

func (c *SpanService) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.client.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *SpanService) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
