package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// create empty list of 52 card deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ = i, because its not used
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver of deck
// d = self / this
// bisa dipanggil dengan namaType.print() => d.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deck, deck = 'go' return 2 value
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// reciver
// ambil string dari deck
func (d deck) toString() string {
	// convert d (deck (list)) to single string, separate with comma (,)
	return strings.Join([]string(d), ",")
}

// error = callback / return type
func (d deck) saveToFile(filename string) error {
	// 0666 file permission => everyone can write & read
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// bs = byte string
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// file tidak error tapi some happen
		// 1. log error dan return a call to newDeck()
		// 2. log error dan kill apps
		fmt.Println("Error: ", err)
		// 0 = progam is ok, else = not ok
		os.Exit(1)
	}

	// convert byte string to string
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// receiver, jadi nanti bisa dipake. deck.shuffle()
func (d deck) shuffle() {
	// source = seed untuk rand
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// random nomor antara 0 - 1
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
