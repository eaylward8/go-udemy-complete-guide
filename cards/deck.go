package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	// return 2 decks
	// 0 to handSize (not inclusive)
	// handSize to end of deck
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// https://golang.org/pkg/io/ioutil/#WriteFile
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// no receiver, we don't yet have a deck, but want one
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // 0 means success, anything else means error
	}

	// string(bs) converts byte slice to string
	// split the string on "," to return a slice of strings
	s := strings.Split(string(bs), ",")
	// use type conversion to convert slice of strings to deck
	// easily possible because a deck is essentially a slice of strings with
	// extra functions attached
	return deck(s)
}

// not returning anything, just shuffling the receiving deck
func (d deck) shuffle() {
	// produce a new "Source" with a different seed based on time stamp
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap elements in deck
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
