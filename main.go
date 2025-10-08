package main

import "fmt"

func main() {

	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Four of Hearts")

	// fmt.Println(cards)

	cards := newDeck()
	// cards.print()

	hand, remainingCards := deal(cards, 3)

	hand.print()

	fmt.Println("Remaining cards---------")

	remainingCards.print()
}
