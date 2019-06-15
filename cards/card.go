package main

import "fmt"

func shuffle() {
	deck := newDeck()
	deck.shuffle()
	deck.print()

}

func withFile() {
	deck := newDeck()
	deck.saveToFile("my_card.txt")
	newDeck := newDeckFromFile("my_card.txt")
	newDeck.print()
}

func handing() {
	deck := newDeck()
	hand, remaingDeck := deck.hand(5)
	print("Handing deck: ========> \n")
	hand.print()
	print("Remaining deck: =========> \n")
	remaingDeck.print()

}

func hello() {
	var afs string = "Ace of Spades"
	fmt.Println(afs)

	fod := "Five of Diamonds"
	fmt.Println(fod)

	var testInt int
	testInt = 10

	fmt.Println(testInt)

	test = 100
	fmt.Println(test)

	x := newCard()
	fmt.Println(x)

	fmt.Println(testCross())

	cards := []string{newCard(), "Ace of Hearts"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	decks := deck{newCard(), newCard()}
	for i, deck := range decks {
		fmt.Println(i, deck)
	}

	decks.print()
}

func newCard() string {
	return "Ten of Clubs"
}

func getPi() float64 {
	return 3.14
}
