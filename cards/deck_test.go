package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "One of Spades" {
		t.Errorf("Expected first card of One of Spades, but got %v", d[0])
	}

	if d[len(d)-1 ] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Deck sizes do not match. %v != %v", len(loadedDeck), len(deck))
	}

	os.Remove("_decktesting")

}