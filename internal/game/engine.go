package game

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func StartGame() {
	players := CreatePlayers()
	current := 0
	playersWon := 0
	gameOver := false

	for !gameOver {
		fmt.Printf("%s's turn\n", players[current].Name)
		fmt.Printf("Your cards:\n")

		ShowDeck(players[current].Hand)

		picks, err := PickCards()

		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println()
			continue
		}

		fmt.Println("You picked:", picks)

		sort.Sort(sort.Reverse(sort.IntSlice(picks)))
		for _, i := range picks {
			if i >= 0 && i < len(players[current].Hand) {
				players[current].Hand = RemoveAtIndex(players[current].Hand, i)
			}
		}

		fmt.Println("Hand length:", len(players[current].Hand))

		if len(players[current].Hand) == 0 {
			fmt.Printf("%s wins!\n", players[current].Name)
			playersWon++
		}

		if playersWon == 3 {
			gameOver = true
		}

		current = (current + 1) % len(players)

		WaitForEnter()
	}

}

func PickCards() ([]int, error) {
	fmt.Print("Pick your cards: ")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	var picks []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid input: %s", part)
		}
		picks = append(picks, num-1)
	}

	return picks, nil
}

func RemoveAtIndex(hand []Card, index int) []Card {
	return append(hand[:index], hand[index+1:]...)
}

func WaitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	_, _ = reader.ReadString('\n')
}
