package main

func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("---")
	// remainingCards.print()

	// str := cards.toString()
	// fmt.Println(str)

	// cards.saveToFile("my_cards")
	// d := newDeckFromFile("my_cards")
	// d.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
