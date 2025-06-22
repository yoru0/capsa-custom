package game

import "fmt"

var playerOneCards, playerTwoCards, playerThreeCards, playerFourCards []Card

func PlayerCards() ([]Card, []Card, []Card, []Card) {
	playerOneCards, playerTwoCards, playerThreeCards, playerFourCards = DealCards()

	fmt.Println("Player 1 Cards:", playerOneCards)
	fmt.Println("Player 1 Cards:", playerTwoCards)
	fmt.Println("Player 1 Cards:", playerThreeCards)
	fmt.Println("Player 1 Cards:", playerFourCards)

	return playerOneCards, playerTwoCards, playerThreeCards, playerFourCards
}