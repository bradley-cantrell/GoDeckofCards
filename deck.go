package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	for i, card := range d { //for loop example
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d { //for loop example
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] //swaps the elements at i and newPosition
	}
}
