package game

import "fmt"

func ShowDeck(deck []Card) {
	for i, card := range deck {
		colorCode := ""

		switch card.Suit {
		case "♥", "♦":
			colorCode = "\033[31m" // red
		case "♠", "♣":
			colorCode = "\033[36m" // cyan
		}

		reset := "\033[0m" // reset

		fmt.Printf("%2d: %s%s%s  ", i+1, colorCode, card.Value+card.Suit, reset)
	}
	fmt.Println()
}
