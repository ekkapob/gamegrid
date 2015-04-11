package main

type Config struct {
	GridRowColumn  string
	MaxRow         int
	MaxColumn      int
	AllowEmptyGame bool
	Games          []GameConfig
}

type GameConfig struct {
	Name  string
	Score string
}
