package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// receiver english bot. func name = getGreeting. return type = string
func (englishBot) getGreeting() string {
	// custom logic for generating english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// custom logic for generating spanish greeting
	return "Hola!"
}
