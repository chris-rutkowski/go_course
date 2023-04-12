package main

import "fmt"

func main() {
	// equivalents:
	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"
	// card := "Ace of Spades"
	// fmt.Println(cards)

	cards := newDeck()
	cards.print()

	fmt.Println("---")

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("---")
	remainingDeck.print()
	fmt.Println(remainingDeck.toString())
	remainingDeck.saveToFile("/Users/chris/deck.txt")
	fmt.Println("---")
	deckFromFile := newDeckFromFile("/Users/chris/deck.txt")
	deckFromFile.shuffle()
	deckFromFile.print()

	// greeting := "Hi there!🇬🇧"
	// fmt.Println([]byte(greeting))
}
