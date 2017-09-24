package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	// hand, remaining := deal(cards, 5)

	// hand.print()
	// remaining.print()

	fmt.Println(cards.toString())
}
