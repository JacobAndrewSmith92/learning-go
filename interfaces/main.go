package main

import "fmt"

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

// in a struct you can remove the first param of the receiver if not being used
// ex. (eb englishBot) -> (englishBot)

func (englishBot) getGreeting() string {
	// VERY CUSTOM LOGIC
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// VERY CUSTOM LOGIC
	return "Hola!"
}
