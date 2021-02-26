package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
//  A slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cloves"}
	cardValues := []string{"Ace", "Two", "There", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
