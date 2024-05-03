package main

import (
	"math/rand"
	"time"
)

type card struct {
	num   int
	color string
}
type hand struct {
	cards [3]card
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
	// Shuffle the deck
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	return deck
}
