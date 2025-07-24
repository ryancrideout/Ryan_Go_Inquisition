package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	// "github.com/chromedp/chromedp"
)

func main() {
	// URL to request
	hostURL := "http://go_react:5173/"
	for i := 1; i < 11; i++ {
		urlToVisit := fmt.Sprintf("%sendpoint%s", hostURL, strconv.Itoa(i))
		//fmt.Println(urlToVisit)

		// Simple GET request:
		// response, err := http.Get(urlToVisit)
		// if err != nil {
		// 	panic(err)
		// }
		// defer response.Body.Close()
		// fmt.Println(response.StatusCode) // 404?

		client := &http.Client{}
		request, _ := http.NewRequest("GET", urlToVisit, nil)
		request.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Go-http-client/1.1)")
		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

		response, err := client.Do(request)

		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))
	}

	// // "github.com/chromedp/chromedp"

	// // Create a Chrome instance
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()

	// // Create a timeout
	// ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	// defer cancel()

	// // Navigate to the page and wait for it to load
	// var htmlContent string
	// err := chromedp.Run(ctx,
	// 	chromedp.Navigate(hostURL),
	// 	// Wait for the React app to render (adjust selector as needed)
	// 	chromedp.WaitVisible("#root > *"),
	// 	// Get the page content after JavaScript execution
	// 	chromedp.OuterHTML("html", &htmlContent),
	// )

	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// 	return
	// }

	// fmt.Println(htmlContent)
}
