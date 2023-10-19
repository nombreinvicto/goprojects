package main

import "fmt"

// Write a function longest() that takes
// an arbitrary number of strings and prints
// the length of the longest one.
func longest(strings ...string) int {
	prevLength := 0
	for _, str := range strings {
		if len(str) > prevLength {
			prevLength = len(str)
		}
	}
	return prevLength
}

func main() {
	fmt.Println(longest("Six", "sleek", "swans",
		"swam", "swiftly", "southwards"))
	fmt.Println(longest("Your", "word", "list", "here"))
}
