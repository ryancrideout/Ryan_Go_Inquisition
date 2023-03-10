package main

import (
	"fmt"
	"strings"
)

func main() {
	var greeting = "Hello world!"

	// Looking at String Length
	fmt.Printf("String Length is: ")
	fmt.Println(len(greeting))

	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}

	fmt.Printf("\n")
	const sample_text = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	/*q flag escapes unprintable characters, with + flag it escapses non-ascii
	  characters as well to make output unambigous */
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sample_text)
	fmt.Printf("\n")

	new_greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(new_greetings, " "))
}
