package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type gameEvents struct {
	Data struct {
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

func GameEvents(time time.Time, gameId string) gameEvents {
	gameEvents := gameEvents{}

	out, err := os.Create("game_events.json")
	if err != nil {
		fmt.Println("Error creating file: " + err.Error())
	}
	defer out.Close()

	resp, err := http.Get(GameEventsURL(time, gameId))
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

func GameEventsURL(time time.Time, gameId string) string {
	gid := gameId
	gid = "gid_" + gid
	gid = strings.Replace(gid, "/", "_", -1)
	gid = strings.Replace(gid, "-", "_", -1)

	host := "http://gd2.mlb.com"
	main := "components/game/mlb"
	year := "year_" + time.Format("2006")
	month := "month_" + time.Format("01")
	day := "day_" + time.Format("02")
	file := "game_events.json"
	return host + "/" + main + "/" + year + "/" + month + "/" + day + "/" + gid + "/" + file
}
