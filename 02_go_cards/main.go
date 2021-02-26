package main

import "fmt"

func main() {
	// Declaring and assigng a string variable
	//  Var is keyword to declare variables
	//  Go is a static typed language so you need to say what time a variable
	// var card string = "Ace of Spades"

	//  A short hand way of declaring and assigning new variables where the type is inferred by Go
	card := "Ace of Spades"

	// This how you assign values to existing variables with = sign
	card = "Five of Diamonds"

	fmt.Println(card)
}
