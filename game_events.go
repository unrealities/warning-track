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
			InningRaw json.RawMessage `json:"inning"`

			Innings []inning
		}
	}
}

type inning struct {
	Num    string     `json:"num"`
	Bottom halfInning `json:"bottom"`
	Top    halfInning `json:"top"`
}

type halfInning struct {
	AtBatRaw json.RawMessage `json:"atbat"`

	AtBats []atBat
}

type atBat struct {
	Home_Team_Runs string `json:"home_team_runs"`
	Away_Team_Runs string `json:"away_team_runs"`
	O              string `json:"o"`
	B1             string `json:"b1"`
	B2             string `json:"b2"`
	B3             string `json:"b3"`
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

	innings := []inning{}
	inn := inning{}
	err = json.Unmarshal(gameEvents.Data.Game.InningRaw, &innings)
	if err == nil {
		gameEvents.Data.Game.Innings = innings
	} else {
		err = json.Unmarshal(gameEvents.Data.Game.InningRaw, &inn)
		if err == nil {
			innings = append(innings, inn)
		}
		gameEvents.Data.Game.Innings = innings
	}

	for k, i := range gameEvents.Data.Game.Innings {
		atBats := []atBat{}
		ab := atBat{}
		err = json.Unmarshal(i.Bottom.AtBatRaw, &atBats)
		if err == nil {
			i.Bottom.AtBats = atBats
		} else {
			err = json.Unmarshal(i.Bottom.AtBatRaw, &ab)
			if err == nil {
				atBats = append(atBats, ab)
				i.Bottom.AtBats = atBats
			} else {
				i.Bottom = halfInning{}
			}
		}

		atBats = []atBat{}
		ab = atBat{}
		err = json.Unmarshal(i.Top.AtBatRaw, &atBats)
		if err == nil {
			i.Top.AtBats = atBats
		} else {
			err = json.Unmarshal(i.Top.AtBatRaw, &ab)
			if err == nil {
				atBats = append(atBats, ab)
				i.Top.AtBats = atBats
			} else {
				i.Top = halfInning{}
			}
		}

		gameEvents.Data.Game.Innings[k] = i
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
