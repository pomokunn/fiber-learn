package main

import (
	"fmt"

	"github.com/pomokunn/hello/exercise"
)

func main() {
	book := exercise.Book{Name: "Harry Potter", Author: "J. K. Rowling"}
	fmt.Println(book.String()) // Output: Harry Potter by J. K. Rowling

	book.SetName("Harry Potter and the Goblet of Fire")
	fmt.Println(book.String()) // Output: Harry Potter and the Goblet of Fire by J. K. Rowling
}
