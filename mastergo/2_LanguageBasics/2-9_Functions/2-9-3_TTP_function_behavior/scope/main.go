package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "abcde"
	for _, s2 := range s1 {
		s3 := unicode.ToUpper(s2)
		fmt.Print(string(s3))
	}
	fmt.Println("\n" + s1)
}
