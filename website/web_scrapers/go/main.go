package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Create a custom HTTP client with longer timeout and that follows redirects
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// URL to request
	hostURL := "http://go_react:5173/"

	// Create request
	req, err := http.NewRequest("GET", hostURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Add headers that curl might be using
	req.Header.Set("User-Agent", "Go-http-client/1.1")
	req.Header.Set("Accept", "*/*")

	// Make the request
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Println("\nResponse Body:")
	fmt.Println(string(body))
}
