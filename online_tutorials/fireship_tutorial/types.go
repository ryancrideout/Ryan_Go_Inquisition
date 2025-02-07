package main

import "fmt"

func main() {
	// User specified types
	const a int32 = 12
	const b float32 = 20.5
	var c complex128 = 1 + 4i
	// uint stands for unsigned integer. In short, unsigned integers
	// can only be positive, but that doesn't paint the whole picture.
	var d uint16 = 14

	// Default Types
	n := 42             // integer
	pi := 3.14          // float64
	x, y := true, false // bool
	z := "Go me! :)"    // string

	fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}
