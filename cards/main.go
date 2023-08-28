package main

func main() {

	// declaring a slice of strings
	cards := newDeck()

	// get a deal
	hand, remainingCards := deal(cards, 5)

	// print contents
	hand.print()
	remainingCards.print()
}
