package main

import (
	"flag"
	"fmt"
)

func main() {
	verbosePtr := flag.Bool("v", false, "to print last name or not")
	firstNamePtr := flag.String("first", "", "first name")
	lastNamePtr := flag.String("last", "", "last name")
	agePtr := flag.Int64("age", 0, "age")

	// parse the comm line
	flag.Parse()
	if *verbosePtr == true {
		fmt.Println(*verbosePtr, *firstNamePtr, *lastNamePtr, *agePtr)
	} else {
		fmt.Println("nothing to print")
		panic("some error")
	}
}
