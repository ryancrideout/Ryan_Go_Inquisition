package main

import "fmt"

func count_to_ten(goroutine_id int) {
	for i := 0; i < 10; i++ {
		fmt.Println(goroutine_id, ":", i)
	}

}

func main() {

}
