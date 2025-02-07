package main

import "fmt"

func main() {
	/* n is an array of 10 integers */
	var n [10]int
	var i, j int

	/* initialize elements of array n to 0 */
	for i = 0; i < 10; i++ {
		/* set element at location i to i + 100 */
		n[i] = i + 100
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

/*
More Array resources
https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm
https://www.tutorialspoint.com/go/go_passing_arrays_to_functions.htm
*/
