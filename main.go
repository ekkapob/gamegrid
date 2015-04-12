package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func loadConfig(url string) (config Config, err error) {
	file, e := ioutil.ReadFile(url)
	if e != nil {
		return config, errors.New(fmt.Sprint("Error: can't load configuration file ", url))
	}
	json.Unmarshal(file, &config)
	return config, nil
}

func main() {
	config, err := loadConfig("config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	board := Board{config: config}
	board.init()

	for {
		board.drawBoard()
		command := board.waitForCommand()
		if strings.ToLower(command) == "x" {
			return
		}
		board.executeCommand(command)
	}
}
