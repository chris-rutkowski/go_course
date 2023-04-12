package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	triangle := triangle{
		base:   2,
		height: 5,
	}
	printArea(triangle)

	square := square{
		sideLength: 3,
	}
	printArea(square)
}

// we can just say englishBot instead of eb enblishBot cuz we don't use eb variable in implementation
// func (englishBot) getGreeting() string {
// func (eb englishBot) getGreeting() string {

// conforming both bots to interface just by fulfilling the interface (all method and correct signatures)
func (englishBot) getGreeting() string {
	return "Hi there!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
