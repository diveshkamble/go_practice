package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expected_length := 52

	if len(d) != expected_length {
		t.Errorf("Expected deck length of %v, but got  %v", expected_length, len(d))
	}
}
