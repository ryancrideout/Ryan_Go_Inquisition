package main

import "fmt"

/*
func main() {
	// Define an array of size 4 that stores deployment options.
	var deployment_options = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	for i := 0; i < len(deployment_options); i++ {
		option := deployment_options[i]
		fmt.Println(i, option)
	}
}
*/

// This is the CLEANER way of doing things:
func main() {
	// Define an array and let the compiler count the size.
	deployment_options := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	// Loop through the deployment options aray
	for index, option := range deployment_options {
		fmt.Println(index, option)
	}
}
