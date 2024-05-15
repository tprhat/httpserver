package main

import (
	"testing"
)

func TestGameCreation(t *testing.T) {
	game := NewGame("Kralj", "Karla")
	if len(game.deck) != 34 {
		t.Errorf("There are not enough cards in the deck! Expected: 34 Received: %d", len(game.deck))
	}
	if len(game.player1.hand.cards) != 3 {
		t.Errorf("Player 1 has an incorrect amount of cardS! Expected: 3 Received: %d", len(game.player1.hand.cards))
	}
	if len(game.player2.hand.cards) != 3 {
		t.Errorf("Player 2 has an incorrect amount of cardS! Expected: 3 Received: %d", len(game.player2.hand.cards))
	}
	if game.main_card.num == 0 {
		t.Errorf("Main card not initialized properly!")
	}
}
