package main

import (
	"os"
	"testing"
)

// command: go test
// if problem: go: go.mod file not found in current directory or any parent directory; see 'go help modules'
// than solution: go env -w GO111MODULE=auto

func TestNewDeck(t *testing.T) {
	sut := newDeck()

	if len(sut) != 20 {
		t.Errorf("Expected deck length of 20 but got %d", len(sut)) // or %v (value)
	}

	if sut[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", sut[0])
	}

	if sut[len(sut)-1] != "Five of Clubs" {
		t.Errorf("Expected last card to be Five of Clubs, but got %v", sut[len(sut)-1])
	}
}

func TestSaveToFileAndLoad(t *testing.T) {
	os.Remove("_decktesting")

	sut := newDeck()
	sut.saveToFile("_decktesting")

	loadedSut := newDeckFromFile("_decktesting")

	if len(loadedSut) != 20 {
		t.Errorf("Expected loaded deck length of 20 but got %d", len(loadedSut))
	}

	os.Remove("_decktesting")
}
