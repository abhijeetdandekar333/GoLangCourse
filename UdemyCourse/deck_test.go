package main

import (
	"testing"
	"os"
)

// the test function should have first letter as capital


func TestNewDeck(t *testing.T)  {
	d := newDeck()
// 1. Test the Code to ensure that a deck is created with x number of cards.
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}
// 2. Test the Code to ensure that first card in deck is "Ace of Spades".
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades , but got %v", d[0])
	}
// 3. Test the Code to ensure that last card in deck is "Four of Clubs".
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}
// 4. Test the Code to ensure that file is saved to local machine and attempt to load file again to create deck.
// Testing saveToDeck and newDeckFromFile functions:

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T)  {
	os.Remove("_decktesting") //Remove the _decktesting file if exists
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 20 {
		t.Errorf("Expected 20 cards, and got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}