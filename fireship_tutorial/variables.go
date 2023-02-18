package main

import "fmt"

// Declare a single variable.
// Variable, Variable Name, Variable Type
var a int

// Declare a GROUP of variables!
// Wait, no commas? Ooooh Jezus
var (
	b bool
	c float32
	d string
)

func main() {
	// We declare all of them variables here.
	a = 42
	b, c = true, 32.0
	d = "string"
	fmt.Println(a, b, c, d)
}
