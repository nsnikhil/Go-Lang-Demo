package main

func main() {
	//cards := deck{"Ace of diamonds","King of hearts"}
	//cards = append(cards, "Eight of spades")

	//cards := newDeck()

	//_ = cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//hand, remainingCard := deal(cards, 5)

	//cards.print()
	//fmt.Println(cards.toString())
	//hand.print()
	//remainingCard.print()

	//for i,card := range cards {
	//	fmt.Println(i,card)
	//}
}
