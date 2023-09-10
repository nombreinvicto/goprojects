package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// t is a test handler - more on this later
	d := newDeck()

	if len(d) != 16 {
		// expecting 4x4=16 cards
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}

	// check first card is ace of spades and last is fur of clubs
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// delete any existing file
	os.Remove("_decktesting")

	// create a new deck and save to file
	deck := newDeck()
	deck.saveToFile("_decktesting")

	// load deck from file
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expecting 16 cards in deck but got %v", len(loadedDeck))
	}

	// cleanup
	os.Remove("_decktesting")
}
