// package site provides methods for initiating status checks against websites
package site

// Site models a website to check
type Site struct {

	// URL specifies the HTTP/HTTPS URL to check
	URL string

	// ExpectedStatusCode is the HTTP status code used to determine "OK" status
	// (default is 200)
	ExpectedStatusCode int

	// LastChecked is a timestamp of the last check run
	LastChecked Time
}
