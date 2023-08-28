package main

import "fmt"

// create a new type of deck which is a slice of strings
// THINK: deck is a "subclass" of slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
}

func (d deck) print() {
	// iterate over slice
	for i, card := range d {
		fmt.Println(i, card)
	}
}
