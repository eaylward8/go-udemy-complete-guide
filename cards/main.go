package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	// cards.print()
	hand.print()
	fmt.Println("--------------")
	remainingCards.print()
}
