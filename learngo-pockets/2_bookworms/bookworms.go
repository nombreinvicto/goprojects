package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
// the counts being the counts of books
// found across all bokworms
func booksCount(bookworms []Bookworm) map[Book]uint {
	bookCounter := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			bookCounter[book] = bookCounter[book] + 1
		}
	}
	return bookCounter
}

func sortBooks(books []Book) []Book {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})
	return books
}

// findCommonBooks returns books that are on more than one bookworm's shelf
func findCommonBooks(bookworms []Bookworm) []Book {
	var commonBooks []Book
	booksOnShelves := booksCount(bookworms)
	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return sortBooks(commonBooks)
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}
