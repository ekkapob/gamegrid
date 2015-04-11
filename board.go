package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	BOARD_ROW    = 8
	BOARD_COLUMN = 8
)

type Board struct {
	row        int
	column     int
	config     Config
	games      [][]Game
	player     Player
	totalScore int
}

func (b *Board) init() {
	// Initialize random source
	rand.Seed(time.Now().UTC().UnixNano())

	b.row = BOARD_ROW
	b.column = BOARD_COLUMN
	b.totalScore = 0

	var row, column int
	if validSize, _ := regexp.MatchString(`^\d+x\d+$`, b.config.GridRowColumn); validSize {
		sizes := strings.Split(b.config.GridRowColumn, "x")
		row, _ = strconv.Atoi(sizes[0])
		column, _ = strconv.Atoi(sizes[1])

		if row > b.config.MaxRow {
			b.row = b.config.MaxRow
		} else if row > 0 {
			b.row = row
		}
		if column > b.config.MaxColumn {
			b.column = b.config.MaxColumn
		} else if column > 0 {
			b.column = column
		}
	}
	b.initializeGames()
	b.initializePlayer()
}

func (b *Board) initializeGames() {
	b.games = make([][]Game, b.row)
	for row, _ := range b.games {
		b.games[row] = make([]Game, b.column)
		for column, _ := range b.games[row] {
			game := getRandomGame(b.config)
			b.games[row][column] = createGame(column, row, game.Name, game.Score)
		}
	}
}

func (b *Board) initializePlayer() {
	b.player.x = rand.Intn(b.column)
	b.player.y = rand.Intn(b.row)
}

func (b *Board) drawBoard() {
	fmt.Println("\nBoard size is", b.row, "x", b.column, "(row x column)")
	playerGamePiece := b.games[b.player.y][b.player.x]
	b.totalScore += playerGamePiece.getScore()

	for row, _ := range b.games {
		fmt.Printf("%5d) ", row+1)

		for column, _ := range b.games[row] {
			playerMarker := " "
			if b.player.y == row && b.player.x == column {
				playerMarker = "*"
			}
			fmt.Print("|", playerMarker, b.games[row][column].getName())
		}
		fmt.Println("|")
	}
	fmt.Println("(*) Player's position is:", b.player.y+1, "x", b.player.x+1, "(row x column)")
	if strings.TrimSpace(playerGamePiece.getName()) == "" && playerGamePiece.getScore() == 0 {
		fmt.Println(`Player: "No game found!"`)
	} else {
		fmt.Println("Player: \"Game", playerGamePiece.getName(), "with score (", playerGamePiece.getScore(), ") found!\"")
	}
	fmt.Println("\nTotal score:", b.totalScore)
	playerGamePiece.clear()
}

func (b Board) waitForCommand() (command string) {
	fmt.Println("Enter command to move, restart, exit")
	fmt.Print("[a]:left, [d]:right, [w]:up, [s]:down, [r]:restart, [x]:exit => ")
	fmt.Scanln(&command)
	return command
}

func (b *Board) executeCommand(command string) {
	switch strings.ToLower(command) {
	case "a":
		if b.player.x > 0 {
			b.player.moveColumn(-1)
		}
	case "d":
		if b.player.x < b.column-1 {
			b.player.moveColumn(1)
		}
	case "w":
		if b.player.y > 0 {
			b.player.moveRow(-1)
		}
	case "s":
		if b.player.y < b.row-1 {
			b.player.moveRow(1)
		}
	case "r":
		b.init()
	}
}
