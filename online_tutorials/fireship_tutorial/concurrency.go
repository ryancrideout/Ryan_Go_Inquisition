package main

import "fmt"

func main() {
	channel := make(chan int) // create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cooking_gopher(i, channel) // Start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopher_id := <-channel // receive a value from a channel
		fmt.Println("gopher", gopher_id, "finished the dish")
	} // All goroutines are finished at this point.
}

// Notice the channel as an argument
func cooking_gopher(id int, channel chan int) {
	fmt.Println("gopher", id, "started cooking!")
	channel <- id // Send a value back to main
}
