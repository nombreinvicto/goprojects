package main

import "fmt"

func main() {

	// declaring a slice of strings
	cards := []string{"Ace of Diamonds", newCard()}

	// add new element to slice
	// append function return NEW slice doesnt modify existing slice
	cards = append(cards, "Six of Spades")

	// iterate over slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
