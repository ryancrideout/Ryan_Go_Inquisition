// Check out this page, pretty fun:
// https://www.tutorialspoint.com/go/go_constants.htm

package main

import "fmt"

func main() {
	// "\t" is tab.
	fmt.Printf("Hello\tWorld!")

	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d", area)
}
