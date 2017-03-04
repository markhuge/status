package site

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstructor(t *testing.T) {
	// test with defaults
	url := "http://foo.bar"
	defaults := New(url)

	if defaults.ExpectedStatusCode != 200 {
		t.Errorf("Defaults: StatusCode expected 200, got %d", defaults.ExpectedStatusCode)
	}

	if defaults.URL != url {
		t.Errorf("Defaults: URL expected %s, got %s", url, defaults.URL)
	}

}

type ResponseFixture struct {
	Body         string
	AlertString  string
	ExpectedCode int
	Code         int
}

var httpResponses = []ResponseFixture{
	{Code: 200, ExpectedCode: 200, Body: "Whatever", AlertString: "OK"},
	{Code: 205, ExpectedCode: 200, Body: "Whatever", AlertString: "WARN"},
	{Code: 500, ExpectedCode: 200, Body: "Whatever", AlertString: "FATAL"},
}

func TestCheck(t *testing.T) {
	for _, fixture := range httpResponses {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(fixture.Code)
			w.Write([]byte(fixture.Body))
		}))
		defer ts.Close()

		s := New(ts.URL)
		response := s.Check()

		if response.StatusCode != fixture.Code {
			t.Errorf("StatusCode expected %d, got %d", fixture.Code, response.StatusCode)
		}

		if response.Body != fixture.Body {
			t.Errorf("StatusCode expected %s, got %s", fixture.Body, response.Body)
		}

		if response.AlertString != fixture.AlertString {
			t.Errorf("AlertString expected %s, got %s", fixture.AlertString, response.AlertString)
		}

	}

}

func TestHTTPError(t *testing.T) {
	s := New("http://unroutable")
	response := s.Check()
	if response.AlertString != "FATAL" {
		t.Errorf("AlertString expected %s, got %s", "FATAL", response.AlertString)
	}

}
