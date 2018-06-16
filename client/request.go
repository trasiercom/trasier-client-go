package client

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func (s *TrasierClient) doRequest(req *http.Request) ([]byte, error) {
	// req.SetBasicAuth(s.Username, s.Password)
	req.Header.Add("Content-Type", "application/json;charset=utf-8");
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
