package menu

import (
	"fmt"
	"os"
)

func Menu() {
	var input int

	for {
		fmt.Println("Welcome to CLI Capsa")
		fmt.Println("1. New Game")
		fmt.Println("2. Exit")
		fmt.Print(">> ")

		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch input {
		case 1:
			fmt.Println("Starting a new game...")
			// TODO: Call game logic here
			continue
		case 2:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input. Please enter 1 or 2.")
		}
	}
}
