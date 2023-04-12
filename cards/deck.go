package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// alias?
type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func newDeckFromFile(path string) deck {
	bytes, error := ioutil.ReadFile(path)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	commaSeparatedCards := string(bytes)
	deckAsString := strings.Split(commaSeparatedCards, ",")
	return deck(deckAsString)
}

// receiver function: deck.print()
// convention: receiver variable - one letter that matches the type deck -> d
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// array[0:3] -> subarray consisting with indexes 0, 1, 2
// array[:2] -> subbarray 0, 1
// array[2:] -> subarray 2, 3, 4 ... count-1
// function returns two values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(path string) error {
	return ioutil.WriteFile(path, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	// source := rand.NewSource(1)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // without seed it would be rand.Intn...
		d[i], d[newPosition] = d[newPosition], d[i] // one line swap (no temp variables)
	}
}
