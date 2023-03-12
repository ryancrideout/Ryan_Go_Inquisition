package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a: %d\n", a)
	fmt.Printf("Before swap, value of b: %d\n", b)

	/*
		Calling a function to swap the values
		&a indicates pointer to a ie. address of varaible a and
		&b indicates pointer to b ie. address of variable b.
	*/
	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	// Save the value at address x
	temp = *x
	// Put y into x
	*x = *y
	// Put temp into y
	*y = temp
}
