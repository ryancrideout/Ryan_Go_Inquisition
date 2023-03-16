package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep for one second
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wait_group sync.WaitGroup

	// So basically the idea here is that we can launch several
	// Goroutines here, but then a waitgroup waits until they're
	// all done before things proceed. Probably worth looking into
	// this some more.
	for i := 1; i <= 5; i++ {
		wait_group.Add(1)

		// Assign the index I guess?
		i := i

		go func() {
			defer wait_group.Done()
			worker(i)
		}()
	}

	wait_group.Wait()
}
