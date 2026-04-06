package main

// entry point, wires everything together
import (
	"flag"
	"fmt"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

// The main function serves as the entry point of the application.
// It loads the configuration from a JSON file, iterates over the list of URLs, launches x amount of goroutines, one for each URL, and prints the results to the console.
func main() {
	// Load configuration from urls.json
	config, err := LoadConfig("urls.json")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	intervalFlag := flag.Int("interval", 60, "Interval in seconds between checks")
	flag.Parse()
	tickerValue := *intervalFlag // Default to 60 seconds if not set as a flag

	// Loop for continuously checking URLs every 60 seconds
	ticker := time.NewTicker(time.Duration(tickerValue) * time.Second)
	defer ticker.Stop()

	for {
		done := make(chan bool, 1)
		go func() {
			for {
				var input string
				fmt.Scanln(&input)
				done <- true
			}
		}()

		for {
			var wg sync.WaitGroup
			// Asynchronously check each URL using goroutines and channels, and use wait.groups to wait for all checks to complete before printing results of each url to the console.
			// Create a buffered channel to collect results from goroutines. The buffer size is set to the number of URLs to prevent blocking.
			// Concurrency limit is set to a set number of goroutines (e.g., 5) to avoid overwhelming the system with too many concurrent requests.
			results := make(chan Result, len(config.URLs))
			sem := make(chan struct{}, 10) // Limit to 10 concurrent goroutines
			for _, url := range config.URLs {
				wg.Add(1)
				go func(u string) {
					defer wg.Done()
					sem <- struct{}{}        // Acquire a slot in the semaphore
					defer func() { <-sem }() // Release the slot after the function completes
					result := CheckURL(u)
					results <- result
				}(url)
			}

			// Wait for all goroutines to finish
			wg.Wait()
			close(results)

			// Collect and print results
			// URL				 | Status | Response Time | Up/Down Status
			// http://example.com: 200    	150ms           UP
			// http://example.com: 403   	200ms           DOWN
			// http://example.com: ---   	---             DOWN (with error message)
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintf(w, "%-30s | %-6s | %-13s | %s\n", "URL", "Status", "Response Time", "Up/Down Status")
			for result := range results {
				fmt.Fprintln(w, result.String())
			}
			w.Flush()

			// --- Countdown + exit ---
			fmt.Printf("\nNext check in %d seconds... Press Enter to exit.\n", tickerValue)

			remaining := tickerValue
			tick := time.NewTicker(1 * time.Second)
			timeout := time.After(time.Duration(tickerValue) * time.Second)

		wait:
			for {
				select {
				case <-done:
					tick.Stop()
					fmt.Println("\nExiting...")
					return
				case <-tick.C:
					remaining--
					fmt.Printf("\rTime remaining: %d seconds ", remaining)
				case <-timeout:
					tick.Stop()
					fmt.Println("\nChecking URLs again...")
					break wait
				}
			}
		}
	}
}
