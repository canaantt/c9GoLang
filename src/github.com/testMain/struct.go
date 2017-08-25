package main

import "fmt"

type Book struct {
	name   string
	author string
	year   int
}

func structUsage() {
	book := Book{name: "Tramp for the Lord", author: "Jamie Cunningham", year: 1976}
	fmt.Println(book)
	fmt.Println(book.name)
}
