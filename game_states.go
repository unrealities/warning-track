package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type gameState struct {
	Inning   int
	Top      bool
	Run_Diff int
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

func GameState(inning int, top bool, run_diff int) int {
	gameStates := GameStates()
	gs := -1

	for key, value := range gameStates {
		if value.Inning == inning && value.Top == top && value.Run_Diff == run_diff {
			gs = key
		}
	}

	return gs
}
