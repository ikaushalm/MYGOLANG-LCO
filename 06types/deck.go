package main

import (
	"fmt"
	"os"
	"strings"
)

//Create a new type of deck which is a slice of string

type deck []string

//we are not going to add reciever as we are creating new
//new deck

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Ace", "Diamond", "Hearts", "Spades"}

	card := []string{"Ace", "Two", "Three", "Four"}
	//_ unused variable
	for _, suit := range cardSuits {
		for _, value := range card {
			cards = append(cards, value+"of"+suit)
		}
	}

	return cards

}

//very clear example of using type
//create a new function belongs to this deck type

// reciever on a function:
// func (d deck ) pring ()
// {} reciver sets up methods the variable type we create

func (d deck) getDeck() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}

//using parameters
// returning multiple values

func deal(d deck, handsize int) (deck, deck) {
	// fmt.Println(d)

	return d[0:handsize], d[handsize:]

}

func (d deck) toByte() []byte {
	newstring := []byte(strings.Join(d, ","))
	return newstring

}

func (d deck) savetoFile(filename string) error {
	// save this deck
	return os.WriteFile(filename, d.toByte(), 0666)

}

func deckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		//we can log the error and close the program
		fmt.Println(err)
		os.Exit(1)
	}

	// string(bs)// Ace of Spades, Two of Spades , Three of Spades

	s := strings.Split(string(bs), ",")
	return deck(s)

}
