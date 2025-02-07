// We need this in order to make this a standalone executable.
package main

// fmt implements formatted I/O.
import "fmt"

/*
When this program is executed, the first function that is run is main.main().

Is it due to package main, then main?
*/
func main() {
	// Call Println() from the fmt package.
	fmt.Println("I am the ember of eternity!")
}
