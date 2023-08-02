package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func Test_newDeck(t *testing.T) {

	t.Run("NewDeck test", func(t *testing.T) {

		const CARD_COUNT int = 52

		d := newDeck()

		if len(d) != CARD_COUNT {
			t.Errorf("count length should be %v, but got %v", CARD_COUNT, len(d))
		}

		if !slices.Contains(d, "Ace of Spades") {
			t.Error("Expected Ace of Spades but it wasn't found.")
		}

	})
}
