package main

type card struct {
	num   int
	color string
}

func buildDeck() []card {
	var deck []card
	nums := []int{1, 2, 3, 4, 5, 6, 7, 11, 12, 13}
	colors := []string{"bastoni", "danari", "spade", "coppe"}
	for _, num := range nums {
		for _, c := range colors {
			deck = append(deck, card{num: num, color: c})
		}
	}
	return deck
}
