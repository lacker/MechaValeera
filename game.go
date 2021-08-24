package main

import "fmt"

type Card int

const (
	COIN Card = iota
	FOXY
	MINSTREL
	PILLAGER
	SCABBS
	SPIRIT
)

type Game struct {
	board []Card // our side of the board
	hand  []Card // our hand
	life  int    // the opponent's life
	mana  int    // our mana
	storm int    // number of things played this turn
}

func NewGame() *Game {
	return &Game{board: []Card{}, hand: []Card{}, life: 30}
}

func main() {
	fmt.Println("hello world")
}
