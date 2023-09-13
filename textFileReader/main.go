package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// first get the filename to open
	filename := os.Args[1]

	// open the file
	filePointer, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

	// reading content from outside(textfile) into aplication
	// so we need Reader implementing type
	io.Copy(os.Stdout, filePointer)

	if err := filePointer.Close(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
	}

}
