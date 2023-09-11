package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// create an empty byteslice
	// make a slice of byte with 99999 empty elements initialised
	// Read function of Reader interface puts data in byte slice
	// until its full. its not designed to grow the byte slice
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

}
