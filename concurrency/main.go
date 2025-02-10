package main

import (
	"fmt"
	"sync"
	"time"
)

func count_to_ten(goroutine_id int, wait_group *sync.WaitGroup) {
	// Once completed, signal to the wait group that this goroutine has completed.
	defer wait_group.Done()
	for i := 0; i < 10000; i++ {
		// fmt.Println(goroutine_id, ":", i)
		continue
		// Add a slight delay so it's more obvious that all of the goroutines are going at once.
		// time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
	}

}

// Build this up so it sends emails as love letters.
// Also demonstrate how we can lock values with multiple channels.
func main() {
	// // Note the "go" variable is what declares a goroutine.
	// go count_to_ten(0)

	start := time.Now()
	var wait_group sync.WaitGroup
	// Multiple goroutines at once!
	for i := 0; i < 100; i++ {
		// Wait groups keep track of goroutines, so we can find out when all goroutines have finished.
		wait_group.Add(1)
		go count_to_ten(i, &wait_group)
	}

	// Wait for all goroutines to finish.
	wait_group.Wait()
	elapsed := time.Since(start)

	// var input string
	// fmt.Scanln(&input)
	fmt.Printf("Total execution time: %v\n", elapsed)

	// We should also time it - how long does it take to do this for 100 goroutines, to count to 10,000?

	// We need to play around with channels here:
	// https://www.golang-book.com/books/intro/10
}
