// Package site provides methods for initiating status checks against websites
package site

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Site models a website to check
type Site struct {

	// URL specifies the HTTP/HTTPS URL to check
	URL string

	// ExpectedStatusCode is the HTTP status code used to determine "OK" status
	// (default is 200)
	ExpectedStatusCode int

	// LastChecked is a timestamp of the last check run
	LastChecked time.Time
}

// New is a constructor function for Site
// accepts a URL string as an argument
func New(url string) *Site {
	s := new(Site)
	s.ExpectedStatusCode = 200
	s.URL = url
	return s
}

// Check performs the HTTP check for the Site
func (s *Site) Check() CheckResponse {

	s.LastChecked = time.Now()

	// start response timer
	start := time.Now()

	// init checkResponse with default values
	checkResponse := CheckResponse{
		URL:         s.URL,
		StatusCode:  0,
		AlertString: "UNKNOWN",
		LastChecked: s.LastChecked,
	}

	res, err := http.Get(s.URL)

	if err != nil {
		checkResponse.Body = err.Error()
		checkResponse.AlertString = "FATAL"
		return checkResponse
	}

	// parse body to string
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	checkResponse.Body = string(bodyBytes)

	checkResponse.StatusCode = res.StatusCode
	// End response timer
	checkResponse.ResponseTime = time.Since(start)
	if res.StatusCode == s.ExpectedStatusCode {
		checkResponse.AlertString = "OK"
		return checkResponse
	}

	if res.StatusCode >= 400 {
		checkResponse.AlertString = "FATAL"
		return checkResponse
	}
	checkResponse.AlertString = "WARN"
	return checkResponse

}
