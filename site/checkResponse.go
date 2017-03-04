package site

import "time"

// CheckResponse represents the consumable output from Check() to be used
// by views/reports
type CheckResponse struct {
	// URL from the request
	URL string

	// StatusCode returned from request
	StatusCode int

	// Body from the request
	Body string

	// LastChecked is the Sites last check time
	LastChecked time.Time

	// ResponseTime of the request
	ResponseTime time.Duration
	// AlertString is either "OK", "WARN", "FATAL"
	AlertString string

	// TODO AlertState int for Nagios-style compatibility
}
