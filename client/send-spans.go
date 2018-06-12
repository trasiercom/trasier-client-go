package client

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"../api"
)

func (s *TrasierClient) SendSpans(spans []api.Span) error {
	url := fmt.Sprintf(baseURL+"/%s/space/%s/spans/", s.AccountId, s.SpaceKey)
	fmt.Println(url)
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
