package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Name struct {
	First string `json:"first_name"`
	Last  string `json:"last_name"`
}

type Book struct {
	Title       string    `json:"book_title"`
	PageCount   int       `json:"pages"`
	ISBN        string    `json:"isbn"`
	Authors     []Name    `json:"authors"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"pub_date"`
}

func main() {
	file, err := os.Open("book.dat")
	if err != nil {
		fmt.Println(err)
		return
	}

	var books []Book
	dec := json.NewDecoder(file)
	if err := dec.Decode(&books); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(books)
}
