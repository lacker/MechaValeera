package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Card int

const (
	COIN Card = iota
	DANCER
	FOXY
	PILLAGER
	SCABBS
	SHARK
)

type CardData struct {
	cost   int
	minion bool
	combo  bool
}

var CardDataMap = map[Card]CardData{
	COIN: {
		cost:   0,
		minion: false,
		combo:  false,
	},
	DANCER: {
		cost:   2,
		minion: true,
		combo:  false,
	},
	FOXY: {
		cost:   2,
		minion: true,
		combo:  false,
	},
	PILLAGER: {
		cost:   6,
		minion: true,
		combo:  true,
	},
	SCABBS: {
		cost:   4,
		minion: true,
		combo:  true,
	},
	SHARK: {
		cost:   4,
		minion: true,
		combo:  false,
	},
}

func (card Card) cost() int {
	return CardDataMap[card].cost
}

func (card Card) minion() bool {
	return CardDataMap[card].minion
}

func (card Card) combo() bool {
	return CardDataMap[card].combo
}

type Game struct {
	board      []Card // our side of the board
	hand       []Card // our hand
	life       int    // the opponent's life
	mana       int    // our mana
	storm      int    // number of things played this turn
	foxy       int    // number of stacks of the foxy effect
	scabbs     int    // number of stacks of the scabbs effect
	nextScabbs int    // number of stacks of the scabbs effect after this one
}

func NewGame() *Game {
	return &Game{board: []Card{}, hand: []Card{}, life: 30}
}

func (game Game) copy() *Game {
	result := &Game{}
	copier.Copy(game, result)
	return result
}

// Mana cost of the card at the given index in hand
// Handles discounts
func (game Game) cost(index int) int {
	card := game.hand[index]
	cost := card.cost()
	cost -= game.scabbs * 3
	if card.combo() {
		cost -= game.foxy * 2
	}
	return cost
}

// Whether we can play the card at the given index in hand
func (game Game) canPlay(index int) bool {
	card := game.hand[index]
	if len(game.board) >= 7 && card.minion() {
		// The board is full
		return false
	}
	return game.mana >= game.cost(index)
}

func (game Game) addToHand(card Card) {
	if len(game.hand) >= 10 {
		return
	}
	game.hand = append(game.hand, card)
}

func (game Game) battlecryAndCombo(card Card) {
	switch card {
	case DANCER:
		game.addToHand(COIN)
	case FOXY:
		game.foxy += 1
	case PILLAGER:
		if game.storm > 0 {
			game.life -= game.storm
		}
	case SCABBS:
		if game.storm > 0 {
			game.scabbs += 1
			game.nextScabbs += 1
		}
	}
}

// Play the card at the given index in hand
func (game *Game) play(index int) {
	card := game.hand[index]
	game.mana -= game.cost(index)
	game.hand = append(game.hand[:index], game.hand[index+1:]...)
	game.foxy = 0
	game.scabbs = game.nextScabbs
	game.nextScabbs = 0

	if card.minion() {
		game.board = append(game.board, card)
	}

	game.battlecryAndCombo(card)
	if game.hasShark() {
		game.battlecryAndCombo(card)
	}

	if card == COIN {
		game.mana += 1
	}

	game.storm += 1
}

func (game *Game) hasShark() bool {
	for _, card := range game.hand {
		if card == SHARK {
			return true
		}
	}
	return false
}

type Move struct {
	index int // which card in hand to play
}

func (game *Game) possibleMoves() []Move {
	answer := []Move{}
	for index := range game.hand {
		if game.canPlay(index) {
			answer = append(answer, Move{index})
		}
	}
	return answer
}

func (game *Game) makeMove(move Move) {
	game.play(move.index)
}

func (game *Game) isWin() bool {
	return game.life <= 0
}

func main() {
	fmt.Println("hello world")
}
