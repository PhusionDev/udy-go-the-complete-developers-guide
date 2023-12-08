package main

func main() {
	cards := newDeck()

	// cards.print()
	hand1, hand2 := cards.deal(5)
	hand1.print()
	hand2.print()
}

func newCard() string {
	return "Five of Diamonds"
}