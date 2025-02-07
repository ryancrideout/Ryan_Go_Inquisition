package main

import "fmt"

/* global variable declaration */
var g int
var global int = 20
var x int = 20

/* function to add two integers */
func sum(x, y int) int {
	// Values passed in have higher precedent than global values.
	fmt.Printf("value of x in sum() = %d\n", x)
	fmt.Printf("value of y in sum() = %d\n", y)

	return x + y
}

func main() {
	/* local variable declaration */
	var a, b, c int
	var global int = 10

	x := 15
	y := 15
	z := 0

	/* actual initialization */
	a = 10
	b = 20
	c = a + b
	g = a * b

	fmt.Printf("value of a = %d, b = %d and c = %d\n", a, b, c)
	fmt.Printf("And g has the value of %d\n", g)
	// I think the takeaway here is that local has higher priority than global.
	fmt.Printf("value of 'global' = %d\n", global)

	fmt.Printf("value of x in main() = %d\n", x)
	z = sum(x, y)
	fmt.Printf("value of z in main() = %d\n", z)
}
