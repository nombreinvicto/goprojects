package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// reading file line by line
func readFile(filename string) {

	// filePointer is *File, File is a struct and is Reader
	filePointer, err := os.Open(filename)
	defer filePointer.Close()
	if err != nil {
		panic(err)
	}

	bytesToRead := 20
	buffer := make([]byte, bytesToRead)

	// infinite loop to go thru file line by line
	for {

		// 0. os.ReadFile reads entire file at a time
		// 1. Read seems to "read" into buffer whatever its length
		bytesRead, err := filePointer.Read(buffer)

		fmt.Println(bytesRead, err)

		// 2. each char/rune is a byte. so if Read is reading
		// 10 bytes at a time, the buffer will contain 10 chars
		// at a time including spaces and escape chars and those chars
		// could be from across different lines in the file
		// Read will populate buffer char by char as it scans
		// across lines in the file

		// towards end of file, buffer could be less than 20 bytes
		// depending on file content left to read. thats why its better
		// to use bytesRead to grab read characters
		//fmt.Println(string(buffer))
		fmt.Println(string(buffer[:bytesRead]))
		fmt.Println(strings.Repeat("=", 50))

		if err == io.EOF {
			break
		}
	}
}
