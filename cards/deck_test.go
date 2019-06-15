package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_desttesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
