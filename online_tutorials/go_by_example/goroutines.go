package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// https://gobyexample.com/goroutines
func main() {
	// Call the function with no threading.
	f("direct")

	// Call the function asynchronously
	go f("goroutine")

	// Call another function asynchronously - this can be interweaved with the
	// other "f" function.
	go func(msg string) {
		fmt.Println(msg)
	}("We are going.")

	// Sleep the function for a secon so everything finished up.
	time.Sleep(time.Second)
	fmt.Println("We are done.")
}
