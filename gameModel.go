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
	gamepiece *GamePiece
}

type GridCell interface {
	getScore() int
}

func (c Cell) getScore() int {
	if c.gamepiece == nil {
		return 0
	}
	if c.gamepiece.Score == "y-pos" {
		return c.column + 1
	}
	if c.gamepiece.Score == "x-pos" {
		return c.row + 1
	}
	score, err := strconv.Atoi(c.gamepiece.Score)
	if err != nil {
		score = 0
	}
	return score
}
