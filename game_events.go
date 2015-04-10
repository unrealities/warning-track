package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type gameEvents struct {
	Subject string
	Data    struct {
		Game struct {
			Inning []inning
		}
	}
}

type inning struct {
	Num    string
	Bottom halfInning
	Top    halfInning
}

type halfInning struct {
	Atbat []atBat
}

type atBat struct {
	Home_Team_Runs string
	Away_Team_Runs string
	O              string
	B1             string
	B2             string
	B3             string
}

func GameEvents() gameEvents {
	gameEvents := gameEvents{}

	gameEventsFile, err := os.Open("example_game_events.json")
	if err != nil {
		fmt.Println("Error opening gameEventsFile: " + err.Error())
	}

	jsonParser := json.NewDecoder(gameEventsFile)
	if err = jsonParser.Decode(&gameEvents); err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return gameEvents
}
