package main

import (
	"encoding/json"
	"os"
)

// Book describes a book on a bookworms shelf
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// Bookworm contains a list of books on a bookworms shelf
type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// loadBookworms reads the file and returns list of bookworms
// and their beloved books found therein
func loadBookworms(filepath string) ([]Bookworm, error) {

	filePointer, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer filePointer.Close()

	// decode the json
	var bookworms []Bookworm

	err = json.NewDecoder(filePointer).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}

// booksCount receives a slice of bookworms
// and returns a map of books and their counts
// the counts being the the counts of books
// found across all bokkworms
func booksCount(bookworms []Bookworm) map[Book]uint {
	bookCounter := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			bookCounter[book] = bookCounter[book] + 1
		}
	}
	return bookCounter
}
