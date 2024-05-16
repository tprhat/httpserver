package main

import (
	"fmt"
	"testing"
)

// func TestTurnError(t *testing.T) {
//     _, err := NewGame("Kralj", "Karla", "Kralj")
//     if err != nil {
//         t.Error(err)
//     }
// }

func TestDeckSize(t *testing.T) {
	game, _ := NewGame("Kralj", "Karla", "Kralj")
	if len(game.deck) != 34 {
		t.Errorf("There are not enough cards in the deck! Expected: 34 Received: %d", len(game.deck))
	}
}

func TestPlayer1HandSize(t *testing.T) {
	game, _ := NewGame("Kralj", "Karla", "Kralj")
	if len(game.player1.hand.cards) != 3 {
		t.Errorf("Player 1 has an incorrect amount of cards! Expected: 3 Received: %d", len(game.player1.hand.cards))
	}
}

func TestPlayer2HandSize(t *testing.T) {
	game, _ := NewGame("Kralj", "Karla", "Kralj")
	if len(game.player2.hand.cards) != 3 {
		t.Errorf("Player 2 has an incorrect amount of cards! Expected: 3 Received: %d", len(game.player2.hand.cards))
	}
}

func TestMainCardInitialization(t *testing.T) {
	game, _ := NewGame("Kralj", "Karla", "Kralj")
	if game.main_card.num == 0 {
		t.Errorf("Main card not initialized properly!")
	}
}

func TestGetCardFromDeck(t *testing.T) {
	game, _ := NewGame("A", "B", "A")
	deck_init_len := len(game.deck)
	game.getCardFromDeck("A")
	if c_type := fmt.Sprintf("%T", game.player1.hand.cards[len(game.player1.hand.cards)-1]); c_type != "main.Card" {
		t.Errorf("c is not a type of Card, it is a type of %s", c_type)
	}
	if deck_init_len-1 != len(game.deck) {
		t.Errorf("Len of deck after removing card is: %d and it should be %d", len(game.deck), deck_init_len-1)
	}
}

func TestRemoveAndAddCardToHand(t *testing.T) {
	game, _ := NewGame("A", "B", "A")
	cards := game.player1.hand.cards
	card := cards[0]
	game.removeCardFromHand(card, "A")
	for _, v := range game.player1.hand.cards {
		if card == v {
			t.Errorf("The card that was player is still in players hand!\nPlayer's hand: %v\nCard played: %v", game.player1.hand.cards, card)
			return
		}
		if len(game.player1.hand.cards) != 2 {
			t.Errorf("Player has the incorrect amount of cards! Expected: 2\tGot:%d", len(game.player1.hand.cards))
		}
	}
}

func TestScoring(t *testing.T) {
	game, _ := NewGame("A", "B", "A")
	game.main_card = Card{num: 3, color: "bastoni"}
	var tests = []struct {
		name  string
		input []Card
		want  Card
	}{
		{"12 bastoni 11 bastoni", []Card{{color: "bastoni", num: 12}, {color: "bastoni", num: 11}}, Card{num: 12, color: "bastoni"}},
		{"12 spade 11 bastoni", []Card{{color: "spade", num: 12}, {color: "bastoni", num: 11}}, Card{num: 11, color: "bastoni"}},
		{"12 coppe 11 denari", []Card{{color: "coppe", num: 12}, {color: "denari", num: 11}}, Card{num: 12, color: "coppe"}},
		{"2 coppe 3 spade", []Card{{color: "coppe", num: 2}, {color: "spade", num: 3}}, Card{num: 2, color: "coppe"}},
		{"2 coppe 3 coppe", []Card{{color: "coppe", num: 2}, {color: "coppe", num: 3}}, Card{num: 3, color: "coppe"}},
		{"3 coppe 2 coppe", []Card{{color: "coppe", num: 3}, {color: "coppe", num: 2}}, Card{num: 3, color: "coppe"}},
		{"1 spade 11 spade", []Card{{color: "spade", num: 1}, {color: "spade", num: 11}}, Card{num: 1, color: "spade"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, _ := game.WinningHand(tt.input[0], tt.input[1])
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestScoreCalculation(t *testing.T) {
	game, _ := NewGame("A", "B", "A")
	game.main_card = Card{num: 3, color: "bastoni"}
	var tests = []struct {
		name  string
		input []Card
		want  int
	}{
		{"12 bastoni 11 bastoni", []Card{{color: "bastoni", num: 12}, {color: "bastoni", num: 11}}, 5},
		{"12 spade 11 bastoni", []Card{{color: "spade", num: 12}, {color: "bastoni", num: 11}}, 5},
		{"12 coppe 11 denari", []Card{{color: "coppe", num: 12}, {color: "denari", num: 11}}, 5},
		{"2 coppe 3 spade", []Card{{color: "coppe", num: 2}, {color: "spade", num: 3}}, 10},
		{"2 coppe 3 coppe", []Card{{color: "coppe", num: 2}, {color: "coppe", num: 3}}, 10},
		{"3 coppe 2 coppe", []Card{{color: "coppe", num: 3}, {color: "coppe", num: 2}}, 10},
		{"1 spade 11 spade", []Card{{color: "spade", num: 1}, {color: "spade", num: 11}}, 13},
		{"7 spade 6 spade", []Card{{color: "spade", num: 7}, {color: "spade", num: 6}}, 0},
		{"1 spade 1 coppe", []Card{{color: "spade", num: 1}, {color: "coppe", num: 1}}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, score := game.WinningHand(tt.input[0], tt.input[1])
			if score != tt.want {
				t.Errorf("got %v, want %v", score, tt.want)
			}
		})
	}
}
