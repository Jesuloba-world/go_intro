package main

import "fmt"

func main() {
	type Book struct {
		name   string
		author string
		year   int
	}

	type Library struct {
		name  string
		books []Book
	}

	myBook := Book{name: "Divine Comedy", author: "Dante Allergeri", year: 1660}

	firstLibrary := Library{name: "First Library", books: []Book{{name: "Dante Allergeri", author: "Dante Allergeri", year: 2009}, {name: "Test book", author: "test author", year: 2016}}}

	fmt.Println(myBook.name)

	fmt.Println(firstLibrary)
}
