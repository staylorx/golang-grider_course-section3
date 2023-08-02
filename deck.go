package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Create a new type of deck which is a slice of strings

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) shuffle() {

	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
