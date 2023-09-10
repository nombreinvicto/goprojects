package main

func main() {
	// read deck from file
	cards := newDeckFromFile("mycards.txt")
	cards.print()

}
