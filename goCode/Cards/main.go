package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// fmt.Println(cards)
	cards = append(cards, "Five of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(cards[0:2])

}

func newCard() string {
	return "Five of Diamonds"
}
