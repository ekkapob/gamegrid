package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func randomGamePiece(games []GamePiece, allowEmptyGame bool) *GamePiece {
	if allowEmptyGame {
		games = append(games, GamePiece{})
	}
	if game := games[rand.Intn(len(games))]; game.Name != "" {
		return &game
	}
	return nil
}

func loadConfig(url string) (config Config, err error) {
	file, e := ioutil.ReadFile(url)
	if e != nil {
		return config, errors.New(fmt.Sprint("Error: can't load configuration file", url))
	}
	json.Unmarshal(file, &config)
	return config, nil
}

func setBoardGame(config Config) (board Board) {
	var row, column int
	if matched, _ := regexp.MatchString(`^\d+x\d+$`, config.GridRowColumn); matched {
		sizes := strings.Split(config.GridRowColumn, "x")
		row, _ = strconv.Atoi(sizes[0])
		column, _ = strconv.Atoi(sizes[1])
	}

	if row > 0 && column > 0 {
		board = Board{row, column}
	} else {
		board = Board{8, 8}
	}

	return board
}

func initializeGamePiece(board Board, config Config) [][]Cell {
	grid := make([][]Cell, board.row)
	for i, _ := range grid {
		grid[i] = make([]Cell, board.column)
		for j, _ := range grid[i] {
			game := &grid[i][j]
			game.row = i
			game.column = j
			game.gp = randomGamePiece(config.Games, config.AllowEmptyGame)
		}
	}
	return grid
}

func drawGrid(grid [][]Cell, player Player) {
	for i, _ := range grid {
		for j, cell := range grid[i] {
			var name, markPlayer string = " ", " "
			if cell.gp != nil {
				name = cell.gp.Name
			}
			if i == player.y && j == player.x {
				markPlayer = "."
			}

			fmt.Print("|", markPlayer, name)
		}
		fmt.Println("|")
	}
}

func getScore(cell GridCell) int {
	return cell.getScore()
}

func getRandomPlayer(board Board) (player Player) {
	playerRow := rand.Intn(board.row)
	playerColumn := rand.Intn(board.column)
	player = Player{playerColumn, playerRow, board}
	return player
}

func clearGamePiece(player Player, grid [][]Cell) {
	grid[player.y][player.x].gp = nil
}

func drawGame(board Board, grid [][]Cell, player Player, totalScore *int) {
	fmt.Println("\n-------->>THE-GRID<<-----------")
	drawGrid(grid, player)
	fmt.Println("Game board's size is", board.row, "x", board.column, "(row x column)")
	fmt.Println("Player's position at row:", player.y+1, "column:", player.x+1, "(see .)")
	*totalScore += getScore(grid[player.y][player.x])
	fmt.Println("Total score:", *totalScore)
}

func listenForCommand() (command string) {
	fmt.Println("Enter command to move, restart, exit")
	fmt.Print("[a]:left, [d]:right, [w]:up, [s]:down, [r]:restart, [x]:exit => ")
	fmt.Scanln(&command)
	return command
}

func main() {
	config, err := loadConfig("config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	board := setBoardGame(config)

	// Generate source for random number
	rand.Seed(time.Now().UTC().UnixNano())
	grid := initializeGamePiece(board, config)

	player := getRandomPlayer(board)
	totalScore := 0

	for {
		drawGame(board, grid, player, &totalScore)
		command := listenForCommand()
		clearGamePiece(player, grid)

		if strings.ToLower(command) == "x" {
			break
		}
		switch strings.ToLower(command) {
		case "a":
			player.moveColumn(-1)
		case "d":
			player.moveColumn(1)
		case "w":
			player.moveRow(-1)
		case "s":
			player.moveRow(1)
		case "r":
			grid = initializeGamePiece(board, config)
			player = getRandomPlayer(board)
			totalScore = 0
		}
	}
}
