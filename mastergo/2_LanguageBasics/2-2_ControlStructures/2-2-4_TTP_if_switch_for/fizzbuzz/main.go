package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {

	// TODO: Implement Fizzbuzz
	i := 1
	for i <= n {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz: ", i)
		case i%5 == 0:
			fmt.Println("Buzz: ", i)
		case i%3 == 0:
			fmt.Println("Fizz: ", i)
		default:
			fmt.Println("Nomatch: ", i)
		}
		i += 1
	}

}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max_, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max_
		}
	}
	fizzbuzz(n)
}
