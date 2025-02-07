package main

import "fmt"

// This puts messages into channels
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// This takes messages from channels and spits them out
func pong(pings <-chan string, pongs chan<- string) {
	// Take message from ping and assign it to "msg"
	msg := <-pings
	// Put a message into pongs
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
