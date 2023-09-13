package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// our own custom type
type logWriter struct{}

func main() {

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// the Write method
func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
