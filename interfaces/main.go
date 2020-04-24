package main

import "fmt"

type bot interface {
	// any types that implement a getGreeting function which returns a string
	// automatically implement the "bot" interface
	getGreeting() string

	// if I uncomment this line, there will be errors because I don't have a bloop function
	// defined for englishBot or spanishBot
	// bloop() int
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// accepts any types that implement the bot interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// can remove receiver variable ("eb") if it's not used in the function body
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
