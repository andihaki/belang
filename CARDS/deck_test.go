package main

import (
	"os"
	"testing"
)

// make sure, new deck berhasil dibuat
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// %v = value pass in
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card is Four of Clubs, but got %v", d[len(d)-1])
	}
}

// test save deck ke file _decktesting
// start & proses = delete file _decktesting
// test passed kalau total data = 16
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
