package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected_length := 52

	if len(d) != expected_length {
		t.Errorf("Expected deck length of %v, but got  %v", expected_length, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {

	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	expected_length := 52

	if len(loadedDeck) != expected_length {
		t.Errorf("Expected deck length of %v, but got  %v", expected_length, len(loadedDeck))
	}

	os.Remove("_decktesting")

}
