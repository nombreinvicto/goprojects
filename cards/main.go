package main

func main() {

	// declaring a slice of strings
	cards := deck{"Ace of Diamonds", newCard()}

	// add new element to slice
	// append function return NEW slice doesnt modify existing slice
	cards = append(cards, "Six of Spades")

	// call the receiver function
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
