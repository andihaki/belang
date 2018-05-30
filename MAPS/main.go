package main

import "fmt"

func main() {
	// [typeDataDariKey]typeDataDariValue
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
