package main

import "fmt"

// https://gobyexample.com/channels
// Probably need to read into channels some more.
func main() {
	// Create a "messages" channel that takes a string input.
	messages := make(chan string)

	// Pass in a "ping" in to the messages channel.
	// Note we've also made this a goroutine
	go func() { messages <- "ping" }()

	// This takes the output of messages and saves it as a message(???)
	msg := <-messages
	fmt.Println(msg)
}
