package game

import "sort"

func IsPair(cards []Card) bool {
	if len(cards) != 2 {
		return false
	}
	return cards[0].Value == cards[1].Value
}

func IsTriple(cards []Card) bool {
	if len(cards) != 3 {
		return false
	}
	return cards[0].Value == cards[1].Value && cards[1].Value == cards[2].Value
}

func IsStraight(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}

	var rank []int
	for _, c := range cards {
		rank = append(rank, ValueRank(c.Value))
	}

	sort.Ints(rank)
	for i := 0; i < len(rank)-1; i++ {
		if rank[i+1] != rank[i]+1 {
			return false
		}
	}

	return true
}

func IsFlush(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}

	for i := 0; i < len(cards)-1; i++ {
		if cards[i].Suit != cards[i+1].Suit {
			return false
		}
	}

	return true
}

func IsFullHouse(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}

	var rank []int
	for _, c := range cards {
		rank = append(rank, ValueRank(c.Value))
	}

	sort.Ints(rank)
	if rank[0] == rank[2] && rank[3] == rank[4] && rank[2] != rank[3] {
		return true
	}
	if rank[0] == rank[1] && rank[2] == rank[4] && rank[1] != rank[2] {
		return true
	}

	return false
}

func IsFourOfAKind(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}

	var rank []int
	for _, c := range cards {
		rank = append(rank, ValueRank(c.Value))
	}

	sort.Ints(rank)
	if rank[0] == rank[3] || rank[1] == rank[4]{
		return true
	}

	return false
}

func IsStraightFlush(cards []Card) bool {
	return IsStraight(cards) && IsFlush(cards)
}
