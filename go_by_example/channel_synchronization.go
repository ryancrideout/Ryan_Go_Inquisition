package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	/*
		I THINK what this does is take boolean channels and then
		shove in "true" values.
	*/
	fmt.Print("Working...")
	// Sleep for ONE second
	time.Sleep(time.Second)
	fmt.Println("Done")

	// Okay I think this sliding a "true" value into a channel that
	// takes Booleans
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// I think this spits out a true...
	// But more importantly, if we removed this line, then
	// the function would exit before this is done.
	// I think this is saying, "Wait until I'm done."
	<-done

}
