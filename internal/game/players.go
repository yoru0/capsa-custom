package game

type Player struct {
	Name string
	Hand []Card
}

func CreatePlayers() []Player {
	p1, p2, p3, p4 := DealCards()

	players := []Player{
		{Name: "Player 1", Hand: p1},
		{Name: "Player 2", Hand: p2},
		{Name: "Player 3", Hand: p3},
		{Name: "Player 4", Hand: p4},
	}

	for i := range players {
		SortCards(players[i].Hand)
	}

	return players
}
