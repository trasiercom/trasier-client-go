package client

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
	"../api"
)

func TestDoRequest(t *testing.T) {
	// given
	spanService := SpanService{AccountId: "TestAccount", SpaceKey: "Testspace"}
	echoHandler := func( w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.FormValue("p"))
	}
	ts := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer ts.Close()

	// when
	err := spanService.SendSpans([]api.Span{}, ts.URL + "/")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
