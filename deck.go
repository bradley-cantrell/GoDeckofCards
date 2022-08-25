package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings, this means deck has all the "features" of a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//if you don't need to use the index variable, you can use an _ instead of something like i
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// (d deck) is a receiver, any variable of type deck now gets access to the print function
// the d is generally based on the type you're using, deck = d, car = c, etc.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card) //note here we actually use i instead of _ because we need to reference it
	}
}

// dealer, (deck, deck) tells go the function will return two values, both of type deck, our hand size is defined in main.go
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// d deck passes the actual deck rather than placing it as an argument in toString()
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
