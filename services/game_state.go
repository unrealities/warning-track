package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/unrealities/warning-track/models"
)

func GameStates() []models.GameState {
	gameStates := []models.GameState{}

	gameStatesFile, err := os.Open("game_state.json")
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

func GameStateToInt(gs string) int {
	switch gs {
	case "Final":
		return 1
	case "Game Over":
		return 2
	case "Postponed":
		return 3
	case "Preview":
		return 11
	case "Pre-Game":
		return 12
	case "Warmup":
		return 13
	case "Delayed":
		return 21
	case "Manager Challenge":
		return 22
	case "In Progress":
		return 23
	}
	return 30
}
