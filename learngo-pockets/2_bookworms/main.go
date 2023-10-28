package main

import (
	"fmt"
	"os"
)

func main() {
	jsonFilePath := "testdata\\bookworms.json"

	bookworms, err := loadBookworms(jsonFilePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)
	fmt.Println("Here are the commonbooks")
	displayBooks(commonBooks)

}
