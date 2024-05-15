package main

import (
	"math/rand"
)

type Game struct {
	deck      []Card
	main_card Card
	player1   player
	player2   player
}

type player struct {
	name   string
	hand   Hand
	points int
	cards_won []Card
}

type Card struct {
	num   int
	color string
}

type Hand struct {
	cards []Card // can I write [3]Card and init it?
}
// pop elements
// x, a = a[len(a)-1], a[:len(a)-1]
func NewGame(name_p1, name_p2 string) Game {
	game := Game{deck: buildDeck(), player1: player{name: name_p1}, player2: player{name: name_p2}}
	// fmt.Printf("%v", game)
	game.main_card = game.deck[0]
	game.player1.hand.cards, game.deck = game.deck[len(game.deck) - 3:], game.deck[:len(game.deck) - 3]	
	game.player2.hand.cards, game.deck = game.deck[len(game.deck) - 3:], game.deck[:len(game.deck) - 3]	

	return game

}
func buildDeck() []Card {
	var deck []Card
	nums := []int{1, 2, 3, 4, 5, 6, 7, 11, 12, 13}
	colors := []string{"bastoni", "danari", "spade", "coppe"}
	for _, num := range nums {
		for _, c := range colors {
			deck = append(deck, Card{num: num, color: c})
		}
	}
	// Shuffle the deck
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	return deck
}

func main() {
	// NewGame()
}
