package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected lenght of 16 but was %v", len(d))
	}

	if d[0] != "ace of spade" {
		t.Errorf("expected first card to be ace of spades but got %v", d[0])
	}

	if d[len(d)-1] != "four of club" {
		t.Errorf("expected last card to be four of club but got %v", d[len(d)-1])

	}

}

//noinspection GoUnhandledErrorResult
func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck_test")

	d := newDeck()
	d.saveToFile("_deck_test")

	loadedDeck := newDeckFromFile("_deck_test")

	if len(loadedDeck) != 16 {
		t.Errorf("expected lenght of 16 but was %v", len(loadedDeck))
	}

	os.Remove("_deck_test")

}
