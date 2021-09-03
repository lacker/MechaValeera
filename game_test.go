package main

import (
	"testing"
)

func TestCanPlay(t *testing.T) {
	game := NewGame()
	game.addCardToHand(COIN)
	if len(game.hand) != 1 {
		t.Fatalf("addToHand did not work")
	}
	if !game.canPlay(0) {
		t.Fatalf("could not play at 0")
	}
}

func TestCanFindLethal(t *testing.T) {
	game := NewGame()
	game.addCardsToHand([]Card{COIN, PILLAGER})
	game.mana = 5
	game.life = 1
	ok, _, _ := game.findWin()
	if !ok {
		t.Fatalf("could not find win")
	}
}
