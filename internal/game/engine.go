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
	winners := make(map[int]bool)
	playersWon := 0
	var winnerList []Winner
	
	
	current := TurnFirst(players)
	fmt.Printf("Player %d's turn.\n", current + 1)
	WaitForEnter()

	for playersWon < 3 {
		if winners[current] {
			current = (current + 1) % len(players)
			continue
		}

		player := &players[current]
		fmt.Printf("%s's turn\n", player.Name)
		fmt.Printf("Your cards [%d card(s)]:\n", len(player.Hand))
		ShowDeck(player.Hand)

		selected, picks, comboName := GetValidCombo(player.Hand)
		fmt.Printf("Combo played: %s\n", comboName)
		fmt.Println("[DEBUG]", selected)

		RemoveSelectedCards(&player.Hand, picks)

		if len(player.Hand) == 0 {
			fmt.Printf("%s wins!\n", player.Name)
			winners[current] = true
			playersWon++
			winnerList = AppendWinner(winnerList, player.Name, playersWon)
		}

		current = (current + 1) % len(players)
		WaitForEnter()
	}
	lastPlaceName := LastPlace(winnerList)
	winnerList = AppendWinner(winnerList, lastPlaceName, 4)

	ShowWinners(winnerList)
}

func RemoveSelectedCards(hand *[]Card, picks []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(picks)))
	for _, i := range picks {
		if i >= 0 && i < len(*hand) {
			*hand = append((*hand)[:i], (*hand)[i+1:]...)
		}
	}
}

func GetValidCombo(hand []Card) ([]Card, []int, string) {
	for {
		picks, err := PickCards(hand)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println()
			continue
		}

		selected := ExtractCards(hand, picks)
		if isValid, combo := CheckCombo(selected); isValid {
			return selected, picks, combo
		}

		fmt.Println("Invalid combo, try again.")
		fmt.Println()
	}
}

func CheckCombo(selected []Card) (bool, string) {
	length := len(selected)

	switch length {
	case 0:
		return true, "Skip"
	case 1:
		return true, "Single"
	case 2:
		if IsPair(selected) {
			return true, "Pair"
		}
	case 3:
		if IsTriple(selected) {
			return true, "Triple"
		}
	case 5:
		switch {
		case IsStraightFlush(selected):
			return true, "Straight Flush"
		case IsFourOfAKind(selected):
			return true, "Four of a Kind"
		case IsFullHouse(selected):
			return true, "Full House"
		case IsFlush(selected):
			return true, "Flush"
		case IsStraight(selected):
			return true, "Straight"
		}
		fmt.Println("Invalid combo...")
	default:
		fmt.Println("Invalid combo length.")
	}

	return false, ""
}

func ExtractCards(hand []Card, picks []int) []Card {
	var selected []Card
	for _, i := range picks {
		if i >= 0 && i < len(hand) {
			selected = append(selected, hand[i])
		}
	}
	return selected
}

func PickCards(hand []Card) ([]int, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Pick your cards: ")
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		var picks []int
		valid := true

		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil || num < 1 || num > len(hand) {
				fmt.Println("Invalid input:", part)
				valid = false
				break
			}
			picks = append(picks, num-1)
		}
		
		if valid {
			return picks, nil
		}
		
		fmt.Println("Please enter valid numbers between 1 and", len(hand))
	}
}

func WaitForEnter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	_, _ = reader.ReadString('\n')
}
