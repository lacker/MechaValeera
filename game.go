package main

import "fmt"

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

// Mana cost of the card at the given index in hand
// Handles discounts
func (game *Game) cost(index int) int {
	card := game.hand[index]
	cost := card.cost()
	cost -= game.scabbs * 3
	if card.combo() {
		cost -= game.foxy * 2
	}
	return cost
}

// Whether we can play the card at the given index in hand
func (game *Game) canPlay(index int) bool {
	card := game.hand[index]
	if len(game.board) >= 7 && card.minion() {
		// The board is full
		return false
	}
	return game.mana >= game.cost(index)
}

// Play the card at the given index in hand
func (game *Game) play(index int) {
	panic("XXX")
}

func main() {
	fmt.Println("hello world")
}
