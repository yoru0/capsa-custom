package game

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Card struct {
	Suit  string
	Value string
}

func CreateDeck() []Card {
	var deck []Card

	for i := 1; i <= 4; i++ {
		var suit string

		switch i {
		case 1:
			suit = "♠"
		case 2:
			suit = "♥"
		case 3:
			suit = "♣"
		case 4:
			suit = "♦"
		}

		for j := 2; j <= 14; j++ {
			var value string

			switch j {
			case 11:
				value = "J"
			case 12:
				value = "Q"
			case 13:
				value = "K"
			case 14:
				value = "A"
			default:
				value = strconv.Itoa(j)
			}

			card := Card{Suit: suit, Value: value}

			deck = append(deck, card)
		}
	}

	return deck
}

func ShowDeck() {
	deck := CreateDeck()
	suitRn := "♠"

	for _, card := range deck {
		if suitRn != card.Suit {
			fmt.Println()
		}

		suitRn = card.Suit

		fmt.Printf("%s%s ", card.Suit, card.Value)
	}

	fmt.Println()
}

func ShuffleDeck(deck []Card) []Card {
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	return deck
}

func DealCards() ([]Card, []Card, []Card, []Card) {
	deck := CreateDeck()
	shuffled := ShuffleDeck(deck)

	p1 := shuffled[:13]
	p2 := shuffled[13:26]
	p3 := shuffled[26:39]
	p4 := shuffled[39:52]

	return p1, p2, p3, p4
}
