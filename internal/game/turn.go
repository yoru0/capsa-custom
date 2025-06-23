package game

import "fmt"

func TurnFirst(players []Player) int {
	var current int
	for i := 0; i < len(players); i++ {
		var picks []int

		fmt.Printf("%s: ", players[i].Name)

		for j := 0; j < len(players[i].Hand); j++ {
			if players[i].Hand[j].Value == "3" {
				fmt.Printf("%s ", players[i].Hand[j].Value+players[i].Hand[j].Suit)

				flag := CheckSpade(players[i].Hand[j].Suit)
				if flag {
					current = i
				}

				picks = append(picks, j)
			}
		}
		RemoveSelectedCards(&players[i].Hand, picks)
		fmt.Println()
	}

	return current
}

func CheckSpade(suit string) bool {
	return suit == "♠"
}
