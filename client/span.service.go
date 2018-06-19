package client

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"../api"
)

type SpanService struct {
	AccountId string
	SpaceKey string
}

func (s *SpanService) SendSpans(spans []api.Span, baseURL string) error {
	url := fmt.Sprintf(baseURL+"accounts/%s/space/%s/spans", s.AccountId, s.SpaceKey)
	j, err := json.Marshal(spans)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

func (s *SpanService) doRequest(req *http.Request) ([]byte, error) {
	// req.SetBasicAuth(s.Username, s.Password)
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
