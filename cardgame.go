package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type Game struct {
	deck      []Card
	main_card Card
	player1   *player
	player2   *player
	turn      int
}

type player struct {
	name      string
	hand      Hand
	points    int
	cards_won []Card
}

type Card struct {
	num   int
	color string
}

type Hand struct {
	cards []Card // can I write [3]Card and init it?
}

var hand_strength = map[int]int{
	1:  10,
	3:  9,
	13: 8,
	12: 7,
	11: 6,
	7:  5,
	6:  4,
	5:  3,
	4:  2,
	2:  1,
}

var scoring = map[int]int{
	1:  11,
	3:  10,
	13: 4,
	12: 3,
	11: 2,
}

// pop elements
// x, a = a[len(a)-1], a[:len(a)-1]
func NewGame(name_p1, name_p2, turn_name string) (*Game, error) {
	var turn int
	if turn_name == name_p1 {
		turn = 0
	} else if turn_name == name_p2 {
		turn = 1
	} else {

		err := fmt.Sprintf("Turn must match either %s or %s!", name_p1, name_p2)
		return nil, errors.New(err)
	}
	game := &Game{
		deck:    buildDeck(),
		player1: &player{name: name_p1, points: 0},
		player2: &player{name: name_p2, points: 0},
		turn:    turn,
	}
	game.main_card = game.deck[0]
	game.player1.hand.cards, game.deck = game.deck[len(game.deck)-3:], game.deck[:len(game.deck)-3]
	game.player2.hand.cards, game.deck = game.deck[len(game.deck)-3:], game.deck[:len(game.deck)-3]

	return game, nil

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
func (g *Game) removeCardFromHand(card Card, player string) {
	var cards []Card
	if player == g.player1.name {
		cards = g.player1.hand.cards
	} else {
		cards = g.player2.hand.cards
	}
	var index int
	for i, v := range cards {
		if v == card {
			index = i
		}
	}
	cards = append(cards[:index], cards[index+1:]...)
	if player == g.player1.name {
		g.player1.hand.cards = cards
	} else {
		g.player2.hand.cards = cards
	}
}

func (g *Game) getCardFromDeck(player string) {
	c := g.deck[len(g.deck)-1]
	g.deck = g.deck[:len(g.deck)-1]
	if player == g.player1.name {
		g.player1.hand.cards = append(g.player1.hand.cards, c)
	} else {
		g.player2.hand.cards = append(g.player2.hand.cards, c)
	}
}

func scoreHand(card1, card2 Card) int {
	return scoring[card1.num] + scoring[card2.num]
}

// card1 is the card from the player who plays first
func (g *Game) WinningHand(card1, card2 Card) (Card, int) {
	score := scoreHand(card1, card2)
	if card1.color == g.main_card.color && card2.color != g.main_card.color {
		return card1, score
	}
	if card1.color != g.main_card.color && card2.color == g.main_card.color {
		return card2, score
	}
	if card1.color == g.main_card.color && card2.color == g.main_card.color {
		if hand_strength[card1.num] > hand_strength[card2.num] {
			return card1, score
		} else {
			return card2, score
		}
	}
	if card1.color == card2.color {
		if hand_strength[card1.num] > hand_strength[card2.num] {
			return card1, score
		} else {
			return card2, score
		}
	}
	return card1, score
}
func main() {
	player1 := "Tomica"
	player2 := "Karla"
	game, err := NewGame(player1, player2, player1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Game Starting!")
	log.Printf("Main card: %v", game.main_card)
	for len(game.player1.hand.cards) > 0 {

		log.Printf("Cards remaining: %d", len(game.deck))
		log.Printf("Player 1 hand: %v", game.player1.hand.cards)
		log.Printf("Player 2 hand: %v", game.player2.hand.cards)

		if game.turn == 0 {
			log.Print("Player 1 plays first")
		} else {
			log.Print("Player 2 plays first")
		}
		var card1, card2 Card
		card1, card2 = game.player1.hand.cards[0], game.player2.hand.cards[0]
		var card_win Card
		var score int
		if game.turn == 0 {
			card_win, score = game.WinningHand(card1, card2)
		} else {
			card_win, score = game.WinningHand(card2, card1)
		}
		if card_win == card1 {
			game.turn = 0
			game.player1.points += score
			game.player1.cards_won = append(game.player1.cards_won, card1, card2)
			log.Printf("Round:\nPlayer1: %v\nPlayer2: %v\nWinner: Player1\nCurrent points P1: %v\tP2: %v",
				card1, card2, game.player1.points, game.player2.points)
			game.removeCardFromHand(card1, player1)
			game.removeCardFromHand(card2, player2)
			if len(game.deck) > 0 {
				game.getCardFromDeck(player1)
				game.getCardFromDeck(player2)
			}
		} else {

			game.turn = 1
			game.player2.points += score
			game.player2.cards_won = append(game.player2.cards_won, card1, card2)

			log.Printf("Round:\nPlayer1: %v\nPlayer2: %v\nWinner: Player2\nCurrent points P1: %v\tP2: %v",
				card1, card2, game.player1.points, game.player2.points)
			game.removeCardFromHand(card1, player1)
			game.removeCardFromHand(card2, player2)
			if len(game.deck) > 0 {
				game.getCardFromDeck(player2)
				game.getCardFromDeck(player1)
			}
		}
		log.Printf("Cards remaining end: %d", len(game.deck))
	}
	if game.player1.points > game.player2.points {
		log.Printf("Winner: Player1\nPoints:%d\nCards won:%v ", game.player1.points, game.player1.cards_won)
	} else if game.player1.points < game.player2.points {
		log.Printf("Winner: Player2\nPoints:%d\nCards won:%v ", game.player2.points, game.player2.cards_won)
	} else {
		log.Println("Tie")
	}
}
