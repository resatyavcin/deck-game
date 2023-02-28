package main

func main() {
	cards := newDeck()
	cards.suffle()
	cards.saveToFile("my_cards")
	cards = newDeckFromFile("my_cards")
	cards.print()
}
