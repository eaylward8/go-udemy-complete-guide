package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type "deck" which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// takes a receiver of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// no receiver
// takes args of deck and int
// returns two decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// https://golang.org/pkg/io/ioutil/#WriteFile
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
