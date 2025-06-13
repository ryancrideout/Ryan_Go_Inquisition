package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// URL to request
	hostURL := "http://go_react:5173/"
	// Create a Chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Navigate to the page and wait for it to load
	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate(hostURL),
		// Wait for the React app to render (adjust selector as needed)
		chromedp.WaitVisible("#root > *"),
		// Get the page content after JavaScript execution
		chromedp.OuterHTML("html", &htmlContent),
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(htmlContent)

	// // Create a custom HTTP client with longer timeout and that follows redirects
	// client := &http.Client{
	// 	Timeout: 10 * time.Second,
	// }

	// // Create request
	// req, err := http.NewRequest("GET", hostURL, nil)
	// if err != nil {
	// 	fmt.Printf("Error creating request: %v\n", err)
	// 	return
	// }

	// // Add headers that curl might be using
	// req.Header.Set("User-Agent", "Go-http-client/1.1")
	// req.Header.Set("Accept", "*/*")

	// // Make the request
	// response, err := client.Do(req)
	// if err != nil {
	// 	fmt.Printf("Error making request: %v\n", err)
	// 	return
	// }
	// defer response.Body.Close()

	// // Read the response body
	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Printf("Error reading response: %v\n", err)
	// 	return
	// }

	// fmt.Println("\nResponse Body:")
	// fmt.Println(string(body))
}
