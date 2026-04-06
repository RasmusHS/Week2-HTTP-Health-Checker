package main

// entry point, wires everything together
import (
	"fmt"
	"sync"
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

	var wg sync.WaitGroup
	// Asynchronously check batches of x URLs using goroutines and channels, and use wait.groups  to wait for all checks to complete before printing results of each url batch to the console.
	batchSize := 5
	results := make(chan Result, len(config.URLs))     // Buffered channel to hold results, with capacity equal to the number of URLs
	for i := 0; i < len(config.URLs); i += batchSize { // Iterate over URLs in batches of x
		end := i + batchSize        // Calculate the end index for the current batch
		if end > len(config.URLs) { // Ensure we don't go out of bounds for the last batch
			end = len(config.URLs) // Adjust end index if it exceeds the total number of URLs
		}
		batch := config.URLs[i:end] // Get the current batch of URLs
		wg.Add(1)                   // Increment the WaitGroup counter for each batch
		go func(urls []string) {    // Launch a goroutine for each batch of URLs
			defer wg.Done()            // Decrement the WaitGroup counter when the goroutine completes
			for _, url := range urls { // Iterate over each URL in the batch
				results <- CheckURL(url) // Send the result of checking the URL to the results channel
			}
		}(batch)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)

	// Collect and print results
	for range config.URLs {
		result := <-results
		fmt.Println(result)
	}
}
