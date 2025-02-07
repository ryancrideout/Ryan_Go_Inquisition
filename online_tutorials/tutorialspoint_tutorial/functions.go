package main

import "fmt"

func max(num_1, num_2 int) int {
	/* local variable declaration */
	var result int

	if num_1 > num_2 {
		result = num_1
	} else {
		result = num_2
	}
	return result
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	ka_pow := max(666, 12)
	fmt.Println(ka_pow)

	a, b := swap("Justice", "Justice_SOUP")
	fmt.Println(a, b)
}
