package spans

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/trasiercom/trasier-client-go/api"
	"log"
)

type SpanService struct {
}

func (s *SpanService) SendSpans(accountId string, spaceKey string, spans []api.Span, baseURL string, authToken string) {
	url := fmt.Sprintf(baseURL+"accounts/%s/space/%s/spans", accountId, spaceKey)
	j, err := json.Marshal(spans)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	req.Header.Set("Authorization", "Bearer "+authToken)

	if err != nil {
		log.Fatal(err)
	}
	_, err = s.doRequest(req)
	if err != nil {
		log.Fatal(err)
	}
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
