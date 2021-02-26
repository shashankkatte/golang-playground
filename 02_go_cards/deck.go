package main

import "fmt"

// Create a new type of 'deck'
//  A slice of strings
type deck []string 

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds","Hearts","Cloves"}
	cardValues := []string {"Ace","Two", "There","Four"}

	for _,suit := range cardSuits {
		for _,value := range cardValues {
			cards = append(cards, value + " of "+ suit)
		}
	}

	return cards
}

func (d deck) print()  {
	for i,card := range d {
		fmt.Println(i, card)
	}
}
