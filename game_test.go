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
