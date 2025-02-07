package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// Declare some Book variables
	var book_1 Books
	var book_2 Books

	// Fill out book attributes
	book_1.title = "I am a cute dog"
	book_1.author = "A cute dog"
	book_1.subject = "Cute dogs"
	book_1.book_id = 333

	book_2.title = "The Joker chose me!"
	book_2.author = "Harvey Dent"
	book_2.subject = "Life sucks"
	book_2.book_id = 666

	/*
		// Now we go ahead and print out all of the information
		fmt.Printf("Book 1 title : %s\n", book_1.title)
		fmt.Printf("Book 1 author : %s\n", book_1.author)
		fmt.Printf("Book 1 subject : %s\n", book_1.subject)
		fmt.Printf("Book 1 book_id : %d\n", book_1.book_id)

		// print Book2 info
		fmt.Printf("Book 2 title : %s\n", book_2.title)
		fmt.Printf("Book 2 author : %s\n", book_2.author)
		fmt.Printf("Book 2 subject : %s\n", book_2.subject)
		fmt.Printf("Book 2 book_id : %d\n", book_2.book_id)
	*/

	// Better way to print out the information!
	printBook(book_1)
	printBook(book_2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
