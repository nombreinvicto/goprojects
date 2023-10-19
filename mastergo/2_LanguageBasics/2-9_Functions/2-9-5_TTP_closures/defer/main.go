package main

import "fmt"

func trace(name string) func() {
	// TODO:
	// 1. Print "Entering <name>"
	fmt.Println("Entering ", name)
	// 2. return a func() that prints "Leaving <name>"
	return func() {
		fmt.Println("Leaving ", name)
	}
}

func f() {
	// TODO: add trace() here so the defer receives the returned function
	defer trace("f")()
	fmt.Println("Doing something")
}

func main() {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}
