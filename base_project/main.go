package main

import (
	"fmt"
	"time"
)

func main() {

	for true {
		fmt.Println("Hello World!")
		d := time.Duration(10e8)
		time.Sleep(d)
	}
}
