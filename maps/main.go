package main

import "fmt"

func main() {
	// multiple ways to declare maps

	// literal syntax
	// map where all keys are string and all values are string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4fbe77",
		"white": "#ffffff",
	}

	// with var keyword (empty map)
	// var colors map[string]string

	// with make keyword
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// colors["black"] = "#000000"
	// delete(colors, "white") // delete key/value pair

	// fmt.Println(colors)
	printMap(colors)
}

// iterating over maps
func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex code for", k, "is", v)
	}
}
