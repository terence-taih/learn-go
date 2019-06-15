package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck', which is a slice of strings

type deck []string // extends and borrow the behaviors of slice of strings

// d: first letter of deck, convertion. Actually, it's "this"
func (d deck) print() {
	for i, deck := range d {
		fmt.Println(i, deck)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// notice that rand.Intn(k) always return the same value because it used the same seed
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) hand(handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func newDeck() deck {
	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	sCards := string(bs)
	deck := deck(strings.Split(sCards, ","))
	return deck
}
