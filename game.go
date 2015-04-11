package main

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type Game interface {
	clear()
	getScore() int
	getName() string
}

type GamePiece struct {
	x         int
	y         int
	name      string
	score     int
	isRemoved bool
}

func (g GamePiece) getName() string {
	return g.name
}

func (g GamePiece) getScore() int {
	return g.score
}

func (g *GamePiece) clear() {
	g.name = " "
	g.score = 0
	g.isRemoved = true
}

type FixedScoreGame struct {
	GamePiece
}

type LocationScoreGame struct {
	GamePiece
}

func (l LocationScoreGame) getScore() int {
	if !l.isRemoved {
		return l.score + 1
	}
	return 0
}

func getRandomGame(config Config) GameConfig {
	if config.AllowEmptyGame {
		config.Games = append(config.Games, GameConfig{" ", "0"})
	}
	return config.Games[rand.Intn(len(config.Games))]
}

func createGame(x, y int, name, gameType string) Game {
	score := 0
	if locationScoreGame, _ := regexp.MatchString(`^[xy]-pos$`, gameType); locationScoreGame {
		score = y
		if strings.HasPrefix(gameType, "x") {
			score = x
		}
		return &LocationScoreGame{GamePiece{x, y, name, score, false}}
	}
	score, _ = strconv.Atoi(gameType)
	return &FixedScoreGame{GamePiece{x, y, name, score, false}}
}
