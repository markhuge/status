package site

import (
	"strconv"
	"time"
)

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

func (c *CheckResponse) String() []string {
	formattedResponseTime := c.ResponseTime - (c.ResponseTime % time.Millisecond)
	return []string{
		c.URL,
		colorStatusCode(c.StatusCode),
		formattedResponseTime.String(),
		c.LastChecked.String(),
	}
}

// TODO this relies on termui, so logic should move into views package
// It's kind of a pain in the ass to do that now, so I'm okay with all the shame.
func colorStatusCode(code int) string {
	codeString := strconv.Itoa(code)

	switch {

	case code >= 200 && code <= 299:
		return "[" + codeString + "](fg-green,bg-default)"

	case code >= 300 && code <= 399:
		return "[" + codeString + "](fg-yellow,bg-default)"

	default:
		return "[" + codeString + "](fg-black,bg-red)"
	}

}
