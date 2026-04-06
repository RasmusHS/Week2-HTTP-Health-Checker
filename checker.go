package main

// HTTP checking logic, goroutines

import (
	"net/http"
	"time"
)

// CheckURL performs an HTTP GET request to the specified URL and returns a Result struct containing the URL, status code, and any error that occurred during the request.
// It uses a timeout of 5 seconds for the HTTP client to prevent hanging on unresponsive URLs.
// Example usage:
// result := CheckURL("http://example.com")
// fmt.Println(result)
func CheckURL(url string) Result {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	start := time.Now()
	resp, err := client.Get(url)
	responseTime := time.Since(start)
	if err != nil {
		return Result{URL: url, Error: err, ResponseTime: responseTime.Round(time.Millisecond)}
	}
	defer resp.Body.Close()
	return Result{URL: url, Status: resp.StatusCode, ResponseTime: responseTime.Round(time.Millisecond)}

}
