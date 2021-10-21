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

func TestCanFindTwoStepLethal(t *testing.T) {
	game := NewGame()
	game.addCardsToHand([]Card{COIN, PILLAGER})
	game.mana = 5
	game.life = 1
	ok, _, _ := game.findWin()
	if !ok {
		t.Fatalf("could not find win")
	}
}

func TestCanFindEightStepLethal(t *testing.T) {
	game := NewGame()
	game.addCardsToHand([]Card{DANCER, FOXY, PILLAGER, PILLAGER, SCABBS, SHARK})
	game.mana = 6
	game.life = 26
	ok, _, _ := game.findWin()
	if !ok {
		t.Fatalf("could not find win")
	}
}

func TestCanFindPotionLethal(t *testing.T) {
	game := NewGame()
	game.addCardsToHand([]Card{SHARK, FOXY, SCABBS, PILLAGER, POTION})
	game.mana = 7
	game.life = 20
	ok, _, _ := game.findWin()
	if !ok {
		t.Fatalf("could not find win")
	}
}

func TestCannotFindNonexistentPotionLethal(t *testing.T) {
	// XXX
}
