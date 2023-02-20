package main

import "fmt"

func main() {
	// Declare an int pointer
	var address *int
	// Declare an int
	number := 42
	// Address stores the memory address of the number
	address = &number
	// Dereferencing the value
	value := *address

	fmt.Printf("address: %v\n", address) // address: 0xc00000e170 - this is different every time!
	fmt.Printf("value: %v\n", value)     // value: 42
}
