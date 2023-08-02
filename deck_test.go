package main

import (
	"os"
	"reflect"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func Test_deck_shuffle(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		{
			name: "example1",
			d:    newDeck(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := tt.d[0]
			tt.d.shuffle()
			newCard := tt.d[0]
			if card == newCard {
				t.Errorf("deck.shuffle() first card is %v, but expected it to be different from original %v.", newCard, card)
			}
		})
	}
}

func Test_deck_toString(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want string
	}{
		{
			name: "example1",
			d:    deck(strings.Split("A-Clubs,5-Diamonds,8-Spades,J-Diamonds,Q-Hearts,7-Clubs,2-Hearts", ",")),
			want: "A-Clubs,5-Diamonds,8-Spades,J-Diamonds,Q-Hearts,7-Clubs,2-Hearts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.toString(); got != tt.want {
				t.Errorf("deck.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deck_saveToFile(t *testing.T) {
	const FILE_NAME = "_decktesting.txt"

	t.Run("File read write testing", func(t *testing.T) {
		err := os.Remove(FILE_NAME)
		if err != nil {
			log.Warn(err)
		}

		deck := newDeck()
		deck.saveToFile(FILE_NAME)

		if got := readFromFile(FILE_NAME); !reflect.DeepEqual(got, deck) {
			t.Errorf("readFromFile() = %v, want %v", got, deck)
		}

		err = os.Remove(FILE_NAME)
		if err != nil {
			log.Warn(err)
		}
	})
}

func Test_deck_deal(t *testing.T) {

	type args struct {
		handSize int
	}
	tests := []struct {
		name              string
		d                 deck
		args              args
		wantDeck          deck
		wantRemainingDeck deck
	}{
		{
			name: "example1",
			d:    deck(strings.Split("A-Clubs,5-Diamonds,8-Spades,J-Diamonds,Q-Hearts,7-Clubs,2-Hearts", ",")),
			args: args{
				handSize: 3,
			},
			wantDeck:          deck{"A-Clubs", "5-Diamonds", "8-Spades"},
			wantRemainingDeck: deck{"J-Diamonds", "Q-Hearts", "7-Clubs", "2-Hearts"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.deal(tt.args.handSize)
			if !reflect.DeepEqual(got, tt.wantDeck) {
				t.Errorf("deck.deal() got = %v, want %v", got, tt.wantDeck)
			}
			if !reflect.DeepEqual(got1, tt.wantRemainingDeck) {
				t.Errorf("deck.deal() got1 = %v, want %v", got1, tt.wantRemainingDeck)
			}
		})
	}
}
