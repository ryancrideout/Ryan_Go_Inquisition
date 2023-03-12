package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	// Take the address of var
	ptr = &a

	// Take the address of ptr using address of operator &
	pptr = &ptr

	// Take the value using pptr
	fmt.Printf("Value of a = %d\n", a)
	fmt.Printf("Value available at *ptr = %d\n", *ptr)
	fmt.Printf("Value available at **ptr = %d\n", **pptr)
}
