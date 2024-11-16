package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck len of 20, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Five of Clubs" {
		t.Errorf("Expected last card of Five of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	filename := "_decktesting"
	d := newDeck()
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck len of 20, but got %v", len(loadedDeck))
	}
}
