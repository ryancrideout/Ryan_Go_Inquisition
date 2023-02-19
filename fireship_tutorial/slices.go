package main

/*
import "fmt"

func main() {
	// Define an array containing programming languages. Must include the trailing comma.
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust",
	}

	// Define the slices.
	classics := languages[0:3]  // alternatively, could do languages[:3]
	modern := make([]string, 4) // the len(modern) is 4
	modern = languages[3:7]     // include 3 and exclude 7
	new := languages[7:9]       // alternatively, could do languages[7:]

	fmt.Printf("classic languages: %v\n", classics) // classic languagues: [C Lisp C++]
	fmt.Printf("modern languages: %v\n", modern)    // modern languages: [Java Python JavaScript Ruby]
	fmt.Printf("new languages: %v\n", new)          // new languages: [Go Rust]
}
*/

// NEW AND IMPROVED SLICES.GO
import (
	"fmt"
	"reflect"
)

func main() {
	// Define an array containing programming languages. Must include the trailing comma.
	languages := [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust",
	}

	all_languages := languages[:]                     // Copy of the array.
	fmt.Println(reflect.TypeOf(all_languages).Kind()) // This is a slice? Quoi?
	fmt.Println(reflect.TypeOf(all_languages))        // This spits out "[]string" instead of "slice"

	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	js_frameworks := frameworks[0:4:4]        // Length 4 (0:4), Capacity 4 (:4)
	frameworks = append(frameworks, "Meteor") // You can't do this with arrays.

	fmt.Printf("all frameworks: %v\n", frameworks)
	fmt.Printf("js frameworks: %v\n", js_frameworks)
}
