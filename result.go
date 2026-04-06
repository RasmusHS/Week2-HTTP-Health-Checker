package main

// result struct and output formatting
import (
	"fmt"
	"time"
)

type Result struct {
	URL          string        `json:"url"`
	Status       int           `json:"status"`
	Error        error         `json:"error,omitempty"`
	ResponseTime time.Duration `json:"response_time,omitempty"`
	UpDownStatus string        `json:"up_down_status,omitempty"`
}

// String method to format the output of the Result struct
// If there is an error, it will include the error message; otherwise, it will show the URL, status code, response time, and up/down status.
// Example output:
// URL				 | Status | Response Time | Up/Down Status
// http://example.com: 200    	150ms           UP
// http://example.com: 403   	200ms           DOWN
// http://example.com: ---   	---             DOWN (with error message)
func (r Result) String() string {
	if r.Error != nil {
		return fmt.Sprintf("%-30s | %-6s | %-13s | DOWN (error: %v)", r.URL, "---", "---", r.Error)
	}

	upDownStatus := "DOWN"
	// Consider status codes in the 200-299 range as "UP"
	// 100 - 199: Informational responses
	// 200 - 299: Successful responses
	// 300 - 399: Redirection messages
	// 400 - 499: Client error responses
	// 500 - 599: Server error responses
	if r.Status >= 200 && r.Status < 300 {
		upDownStatus = "UP"
	}
	if r.Status >= 300 && r.Status < 400 {
		upDownStatus = "REDIRECT"
	}
	if r.Status >= 400 && r.Status < 500 {
		upDownStatus = "DOWN (client error)"
	}
	if r.Status >= 500 && r.Status < 600 {
		upDownStatus = "DOWN (server error)"
	}

	return fmt.Sprintf("%-30s | %-6d | %-13s | %s", r.URL, r.Status, r.ResponseTime, upDownStatus)
}
