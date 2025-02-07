package main

import "fmt"

func main() {
	// Define a map containing the release year of several languages.
	first_releases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	// Look through each entry and output the name and release year.
	for k, v := range first_releases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}

	// NOTE: For Go, dictionaries are always in random order, unlike Python.
}
