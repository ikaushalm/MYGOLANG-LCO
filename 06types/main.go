package main

func main() {
	// cards := newDeck()

	// cards.getDeck()

	// hand, remainingdeck := deal(cards, 3)

	// hand.getDeck()

	// remainingdeck.getDeck()

	// fmt.Println(cards.toByte())

	cards := deckFromFile("fdsff")

	cards.getDeck()

}

func newCard() string {
	return "New Card Added"
}