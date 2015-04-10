package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type gameState struct {
	Inning   int8
	Top      bool
	Run_Diff int8
}

func GameStates() []gameState {
	gameStates := []gameState{}

	gameStatesFile, err := os.Open("game_states.json")
	if err != nil {
		fmt.Println("Error opening gameStatesFile: " + err.Error())
	}

	jsonParser := json.NewDecoder(gameStatesFile)
	if err = jsonParser.Decode(&gameStates); err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return gameStates
}
