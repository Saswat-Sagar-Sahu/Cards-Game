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

	hand, remainingCards := deal(cards, 3)

	hand.print()
	fmt.Println("----------Remaining cards---------")
	remainingCards.print()

	// fmt.Println(hand.toString())

	cards.saveToFile("my_cards")

	cardFromFile := newDeckFromFile("my_cards")
	fmt.Println("----------card from file --------")
	cardFromFile.print()

	cards.shuffle()
	fmt.Println("-------After Shuffle----------")
	cards.print()
}
