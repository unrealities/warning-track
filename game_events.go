package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	out, err := os.Create("game_events.json")
	if err != nil {
		fmt.Println("Error creating file: " + err.Error())
	}
	defer out.Close()

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Error accessing file: " + err.Error())
	}
	defer resp.Body.Close()

	gameEventsFile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading gameEventsFile: " + err.Error())
	}

	err = json.Unmarshal(gameEventsFile, &gameEvents)
	if err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return gameEvents
}
