package main

import "fmt"

// Maps are similar to dicts in Python - they contain a mapping of key-value pairs.

// Both the keys and values are statically typed.

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "##00ff00",
		"white": "#ffffff",
	}
	printMap(colors)
	delete(colors, "white")
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for " + color + " is " + hex)
	}
}
