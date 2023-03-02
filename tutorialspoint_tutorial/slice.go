package main

import "fmt"

func main() {
	/*
		// a slice of unspecified size
		var numbers []int
		numbers == []int{0,0,0,0,0}
	*/

	// A slice of length 3 and capacity of 5. Values are [0, 0, 0]
	// Length is the default length, capacity is how long it can become.
	var numbers = make([]int, 3, 5)
	print_slice(numbers)

	// Nil Slices
	var nil_numbers []int
	print_slice(nil_numbers)

	if nil_numbers == nil {
		fmt.Printf("slice is nil")
	}

	// Subslicing
	power_numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	print_slice(power_numbers)

	// Print original slice
	fmt.Println("power_numbers == ", power_numbers)

	// print the sub slice starting from index 1(included) to index 4(excluded)
	fmt.Println("power_numbers[1:4] == ", power_numbers[1:4])

	// Missing lower bound implies 0
	fmt.Println("power_numbers[:3] == ", power_numbers[:3])

	// Missing upper bound implies len(s)
	fmt.Println("power_numbers[4:] == ", power_numbers[4:])

	numbers_1 := make([]int, 0, 5)
	print_slice(numbers_1)

	// print the sub slice starting from index 0 (included) to index 2 (excluded)
	numbers_2 := power_numbers[:2]
	print_slice(numbers_2)

	// print the sub slice starting from index 2 (included) to index 5 (excluded)
	numbers_3 := power_numbers[2:5]
	print_slice(numbers_3)

	// Append and Copy Functions
	var chump_numbers []int
	print_slice(chump_numbers)

	// Append allows nil slice
	chump_numbers = append(chump_numbers, 0)
	print_slice(chump_numbers)

	// add one element to slice
	chump_numbers = append(chump_numbers, 1)
	print_slice(chump_numbers)

	// Add more than one element at a time.
	chump_numbers = append(chump_numbers, 2, 3, 4)
	print_slice(chump_numbers)

	// create a slice numbers_4 with double the capacity of earlier slice
	numbers_4 := make([]int, len(chump_numbers), (cap(chump_numbers))*2)

	// copy content of numbers to numbers_4
	copy(numbers_4, chump_numbers)
	print_slice(numbers_4)
}

// Helper function!
func print_slice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
