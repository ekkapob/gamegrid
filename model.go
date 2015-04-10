package main

import (
	"strconv"
)

type Config struct {
	GridRowColumn  string
	AllowEmptyGame bool
	Games          []GamePiece
}

type GamePiece struct {
	Name  string
	Score string
}

type Board struct {
	row    int
	column int
}

type Cell struct {
	Board
	gp *GamePiece
}

type GridCell interface {
	getScore() int
}

func (c Cell) getScore() int {
	if c.gp == nil {
		return 0
	}
	if c.gp.Score == "y-pos" {
		return c.column + 1
	}
	if c.gp.Score == "x-pos" {
		return c.row + 1
	}
	score, err := strconv.Atoi(c.gp.Score)
	if err != nil {
		score = 0
	}
	return score
}

type Player struct {
	x     int
	y     int
	board Board
}

func (p *Player) moveColumn(direction int) {
	if direction > 0 && p.x < (p.board.column-1) {
		p.x++
	} else if direction < 0 && p.x > 0 {
		p.x--
	}
}

func (p *Player) moveRow(direction int) {
	if direction > 0 && p.y < (p.board.row-1) {
		p.y++
	} else if direction < 0 && p.y > 0 {
		p.y--
	}
}
