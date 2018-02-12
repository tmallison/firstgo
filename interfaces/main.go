package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Unique logic
	return "Hello"
}

func (spanishBot) getGreeting() string {
	// Unique logic
	return "Hola!"
}