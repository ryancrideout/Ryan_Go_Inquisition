package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// URL to request
	hostURL := "http://go_react:5173/"
	for i := 1; i < 11; i++ {
		urlToVisit := fmt.Sprintf("%sendpoint%s", hostURL, strconv.Itoa(i))

		// Create a Chrome instance
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()

		// Create a timeout
		ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		// Navigate to the page and wait for it to load
		var htmlContent string
		err := chromedp.Run(ctx,
			chromedp.Navigate(urlToVisit),
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
