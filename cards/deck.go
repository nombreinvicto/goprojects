package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create a new type of deck which is a slice of strings
// THINK: deck is a "subclass" of slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// func to print deck
func (d deck) print() {
	// iterate over slice
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func to deal a hand of certain size
func deal(d deck, handSize int) (deck, deck) {
	// slicing slices LOL
	return d[:handSize], d[handSize:]
}

// func to convert deck to single string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// func to save deck to file
func (d deck) saveToFile(filename string) error {
	deckString := d.toString()
	return ioutil.WriteFile(filename, []byte(deckString), 0666)
}

// func to read deck from file
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// byteslice to string to deck
	returnedString := string(byteSlice)
	sliceOfStrings := strings.Split(returnedString, ",")
	return sliceOfStrings

}
