package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	/* Declaration of Book Variables */
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

	// Print book info
	print_book(&book_1)
	print_book(&book_2)
}

func print_book(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
