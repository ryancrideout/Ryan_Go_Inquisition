package main

import "fmt"

func main() {
	/* PHASE ONE OF POINTERS */
	var b int = 10
	fmt.Printf("Address of b variable: %x\n", &b)

	/* PHASE TWO OF POINTERS */
	// Variable declaration
	var a int = 20
	// Pointer declaration
	var ip *int

	// Store address of a in pointer variable
	ip = &a

	fmt.Printf("Address of a variable: %x\n", &a)

	// Address stored in a pointer variable
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	// Access the value using the pointer
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	/* PHASE THREE OF POINTERS */
	// Nil Pointers
	var ptr *int

	fmt.Printf("The value of ptr is : %x\n", ptr)
}

/*
More Pointer References:
https://www.tutorialspoint.com/go/go_array_of_pointers.htm
https://www.tutorialspoint.com/go/go_pointer_to_pointer.htm
https://www.tutorialspoint.com/go/go_passing_pointers_to_functions.htm
*/
