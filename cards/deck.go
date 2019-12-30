package main

import "fmt"

// Create a new type 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// loop through deck and print each value
// d is a variable reference to the deck that received the function call
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this function returns 2 values, both are decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
