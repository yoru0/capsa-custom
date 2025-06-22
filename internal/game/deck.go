package game

import (
	"math/rand"
	"sort"
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

func SuitRank(suit string) int {
	switch suit {
	case "♠":
		return 4
	case "♥":
		return 3
	case "♣":
		return 2
	case "♦":
		return 1
	default:
		return 0
	}
}

func ValueRank(value string) int {
	switch value {
	case "3":
		return 1
	case "4":
		return 2
	case "5":
		return 3
	case "6":
		return 4
	case "7":
		return 5
	case "8":
		return 6
	case "9":
		return 7
	case "10":
		return 8
	case "J":
		return 9
	case "Q":
		return 10
	case "K":
		return 11
	case "A":
		return 12
	case "2":
		return 13
	default:
		return 0
	}
}

func SortCards(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Value != cards[j].Value {
			return ValueRank(cards[i].Value) < ValueRank(cards[j].Value)
		}
		return SuitRank(cards[i].Suit) < SuitRank(cards[j].Suit)
	})
}
