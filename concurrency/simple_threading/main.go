package main

import (
	"fmt"
	"sync"
	"time"
)

func count_to_ten(goroutine_id int, wait_group *sync.WaitGroup) {
	// Once completed, signal to the wait group that this goroutine has completed.
	defer wait_group.Done()
	for i := 0; i < 1000000; i++ {
		// fmt.Println(goroutine_id, ":", i)
		continue
	}

}

func main() {

	start := time.Now()
	var wait_group sync.WaitGroup
	// Multiple goroutines at once!
	for i := 0; i < 1000; i++ {
		// Wait groups keep track of goroutines, so we can find out when all goroutines have finished.
		wait_group.Add(1)
		go count_to_ten(i, &wait_group)
	}

	// Wait for all goroutines to finish.
	wait_group.Wait()
	elapsed := time.Since(start)

	// Average run time: 0.045 - 0.06 seconds
	// Total execution time: 52.2371ms
	// Total execution time: 47.4444ms
	// Total execution time: 52.055ms
	// Total execution time: 60.9834ms
	// Total execution time: 50.8007ms
	// Total execution time: 45.2081ms
	fmt.Printf("Total execution time: %v\n", elapsed)

}
