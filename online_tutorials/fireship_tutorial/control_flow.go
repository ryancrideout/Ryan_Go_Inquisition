package main

import "fmt"

// Regular implementation.
/*
// Define a function to find the average fo floats contained in a slice.
// First define x as the input variable, which is a slice?
// Then define the output, which in this case is "average"? Yes.
func average(slice []float64) (average float64) {
	total := 0.0
	if len(slice) == 0 {
		average = 0
	} else {
		for _, value := range slice {
			total += value
		}
		average = total / float64(len(slice))
	}

	return
}
*/

// Switch Statement implementation.
func average(slice []float64) (average float64) {
	total := 0.0
	// Note we can have variables for switch statements.
	// That's HUGE.
	switch len(slice) {
	case 0:
		average = 0
	default:
		for _, v := range slice {
			total += v
		}
		average = total / float64(len(slice))
	}
	return
}

func main() {
	slice := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(slice)) // 19.197499999999998
}
