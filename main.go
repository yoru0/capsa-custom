package main

import (
	"fmt"

	"github.com/yoru0/capsa-custom/internal/game"
	// "github.com/yoru0/capsa-custom/internal/menu"
)

func main() {
	players := game.CreatePlayers()

	current := TurnFirst(players)

	fmt.Println(current)

	// game.StartGame()

	// menu.Menu()
}

func TurnFirst(players []game.Player) int {
	for i := 0; i < len(players); i++ {
		fmt.Println(players[i].Hand)

		fmt.Print("Value: ")
		for j := 0; j < len(players[i].Hand); j++ {
			fmt.Printf("%s ", players[i].Hand[j].Value)
		}
		fmt.Println()
	}

	return 1
}