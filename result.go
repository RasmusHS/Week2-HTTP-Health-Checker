package main

// result struct and output formatting
import (
	"fmt"
)

type Result struct {
	URL    string `json:"url"`
	Status int    `json:"status"`
	Error  error  `json:"error,omitempty"`
}

// String method to format the output of the Result struct
// If there is an error, it will include the error message; otherwise, it will show the URL and status code.
// Example output:
// http://example.com: 200
// http://example.com: Error - Get "http://example.com": dial tcp: lookup example.com: no such host
func (r Result) String() string {
	if r.Error != nil {
		return fmt.Sprintf("%s: Error - %v", r.URL, r.Error)
	}
	return fmt.Sprintf("%s: %d", r.URL, r.Status)
}
