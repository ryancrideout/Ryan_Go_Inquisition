package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func generateMockData(endpoint string) map[string]interface{} {
	// Generate random number of users (1-10)
	userCount := rand.Intn(10) + 1
	users := make([]map[string]interface{}, userCount)
	for i := 0; i < userCount; i++ {
		users[i] = map[string]interface{}{
			"id":    i + 1,
			"name":  fmt.Sprintf("User%d", i+1),
			"email": fmt.Sprintf("user%d@example.com", i+1),
		}
	}

	// Generate random number of products (1-5)
	productCount := rand.Intn(5) + 1
	products := make([]map[string]interface{}, productCount)
	for i := 0; i < productCount; i++ {
		products[i] = map[string]interface{}{
			"id":    i + 1,
			"name":  fmt.Sprintf("Product %d", i+1),
			"price": rand.Intn(100) + 10,
		}
	}

	return map[string]interface{}{
		"id":          rand.Intn(1000),
		"endpoint":    endpoint,
		"timestamp":   time.Now().Format(time.RFC3339),
		"users":       users,
		"products":    products,
		"randomValue": rand.Float64(),
		"status":      "success",
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	endpoint := request.URL.Path[1:] // This removes the Leading "/"
	if endpoint == "" {
		endpoint = "root"
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(generateMockData(endpoint))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("JSON server running on :8080")
	http.ListenAndServe(":8080", nil)

	// How to test:
	/*
		response, err := http.Get("http://json-server:8080/endpoint1")
		// Returns pure JSON, no JavaScript needed
	*/
}
