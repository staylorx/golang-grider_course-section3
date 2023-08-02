package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	cards := newDeck()

	log.Info("Shuffling the deck...")
	cards.shuffle()

	log.Info("Dealing 5 cards...")
	hand, _ := cards.deal(5)
	hand.print()
	log.Info(hand.toString())

	log.Info("Saving to file...")
	hand.saveToFile("test.txt")

	log.Info("Reading from file...")
	readFromFile("test.txt")
	hand.print()
}

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	// parse string, this is built-in feature of logrus
	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}
	// set global log level
	log.SetLevel(ll)
}

// Create a new deck for all 52 cards.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cardName := fmt.Sprintf("%s of %s", cardValue, cardSuit)
			cards = append(cards, cardName)
		}
	}

	return cards
}
