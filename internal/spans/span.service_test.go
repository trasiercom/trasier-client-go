package spans

import (
	"testing"
	"net/http"
	"fmt"
	"net/http/httptest"
	"github.com/trasiercom/trasier-client-go/api"
)

func TestSendSpans(t *testing.T) {
	// given
	spanService := SpanService{}
	echoHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.FormValue("p"))
	}
	ts := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer ts.Close()

	// when
	err := spanService.SendSpans("awesomeAccount", "awesomeKey", []api.Span{}, ts.URL+"/")
	// then
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestMustReturnHTTPErrors(t *testing.T) {
	// given
	spanService := SpanService{}
	echoHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
	}
	ts := httptest.NewServer(http.HandlerFunc(echoHandler))
	defer ts.Close()

	// when
	err := spanService.SendSpans("awesomeAccount", "awesomeKey", []api.Span{}, ts.URL+"/")
	// then
	if err == nil {
		t.Errorf("Error: %v", err)
	}
}

func TestMustThrowUnsupportedTypeException(t *testing.T) {
	spanService := SpanService{}
	err := spanService.SendSpans("awesomeAccount", "awesomeKey", nil, "")
	if err == nil {
		t.Errorf("Error: no error thrown: Expected Method to throw UnsupportedTypeException")
	}
}
