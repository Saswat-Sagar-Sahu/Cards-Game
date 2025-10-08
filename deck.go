package main

import "fmt"

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
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardVal := range cardValues {
			cards = append(cards, cardVal+" of "+cardSuit)
		}
	}

	return cards
}

//deal with cards
func deal(d deck, handSize int) (deck, deck){

	return d[:handSize], d[handSize:]
}