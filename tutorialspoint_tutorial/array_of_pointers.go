package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{10, 100, 200}

	var i int
	var pointer [MAX]*int

	for i = 0; i < MAX; i++ {
		// Assign the address of the integer
		pointer[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *pointer[i])
	}
}
