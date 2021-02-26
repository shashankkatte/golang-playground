# GO Cards Dealer

This is a simple go app that simulates playing with real deck of playing cards. It has the following main functionalities

## Functionalities

1. **newDeck** Creates a list of playing cards. essentially an array of strings

2. **print** Log out the contents of a deck of cards

3. **Shuffle** Shuffels all the cards in a deck

4. **deal** create a hand of cards

5. **saveToFile** Save a list of cards to file on local storage

6. **newDeckFromFile** Load a list of cards from the local machine

## Concepts Covered

- [X] Variable declarations
- [X] Function and return types
- [X] Slices and For loops
- [ ] OO vs Go Approach
- [ ] Custom Type declarations
- [ ]  Creating a new Dec in our APP
- [ ] Slice Range Syntax
- [ ] Multiple return Values
- [ ] Byte Slices
- [ ] Convert Deck to string
- [ ] Joining Strings
- [ ] Saving data to local storage
- [ ] Reading from Harddrive
- [ ] Error Handling
- [ ] Shuffling Deck
- [ ] Random number generation
- [ ] Testing in Go
- [ ] Writing useful tests
- [ ] Asserting elements in a Slice
- [ ] Testing file IO

## Common data types in Go

### Array

It is a fixed length array, very basic datastructure it cant shrink or grow. Every element should be an identical type

### Slice

This is like an array but can grow and shrink. Evevry element should be an identical type.

## OO vs Go Approach

So normally in an Object oriented language like Java you would have clasees and methods associated. 

However in go there is no concept of Classes. If you need a deck of cards you define Cards

We extend a base type to kinda make it work as class with associated functions.

> Functions with a receiver

