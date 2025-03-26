package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func count_to_ten(goroutine_id int, wait_group *sync.WaitGroup) {
	// Once completed, signal to the wait group that this goroutine has completed.
	defer wait_group.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(goroutine_id, ":", i)
		// Add a slight delay so it's more obvious that all of the goroutines are going at once.
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
	}

}

// Also demonstrate how we can lock values with multiple channels.
func main() {
	// Start the timer
	start := time.Now()

	// Declare a wait group - Essentially a wait group is something that keeps track of goroutines.
	// It WAITS for goroutines to finish.
	var wait_group sync.WaitGroup

	// Multiple goroutines at once!
	for i := 0; i < 10; i++ {
		// Add a tally to the wait group for each iteration of the for loop, one for each goroutine that is spun up.
		wait_group.Add(1)
		// General formula of spinning up a goroutine:
		// go <function>
		go count_to_ten(i, &wait_group)
	}

	// Wait for all goroutines to finish.
	wait_group.Wait()

	// Spit out how long this has taken to run.
	elapsed := time.Since(start)
	fmt.Printf("Total execution time: %v\n", elapsed)
}
