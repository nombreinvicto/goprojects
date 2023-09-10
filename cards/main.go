package main

import "fmt"

func main() {
	// read deck from file
	cards := newDeckFromFile("mycards.txt")
	cards.print()

	fmt.Println("After Shuffle")
	cards.shuffle()
	cards.print()

}
