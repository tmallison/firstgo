package main

import "fmt"

// Create a new type of 'deck'
// wich is a slice of strings
type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heart", "Clubs"}
	cardValue := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValue{
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i+1, card)
	}
}