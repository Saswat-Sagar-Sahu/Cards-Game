package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type of 'deck'
// which is slice of strings
type deck []string

// loop deck of cards & print each element
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// return deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardVal := range cardValues {
			cards = append(cards, cardVal+" of "+cardSuit)
		}
	}

	return cards
}

// deal with cards
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]
}

// utilty to convert deck of card to single string
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	cardsDeck, error := ioutil.ReadFile(fileName)
	if error != nil {
		// option 1 - log the error and return a call to newDeck()
		// option 2 - log the error and entirely quit the program
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	s := strings.Split(string(cardsDeck), ", ")
	return deck(s)
}

// shuffle the deck
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
