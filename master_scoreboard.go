package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

type scoreboard struct {
	Data struct {
		Games struct {
			Game []mlbApiGame `json:"game"`
		}
	}
}

func MasterScoreboard(time time.Time, r *http.Request) scoreboard {
	scoreboards := scoreboard{}

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(MasterScoreBoardURL(time))
	if err != nil {
		fmt.Println("Error accessing MasterScoreBoardURL: " + err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	msbData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading msbData: " + err.Error())
	}

	err = json.Unmarshal(msbData, &scoreboards)
	if err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return scoreboards
}

func MasterScoreBoardURL(time time.Time) string {
	host := "http://gd2.mlb.com"
	main := "components/game/mlb"
	year := "year_" + time.Format("2006")
	month := "month_" + time.Format("01")
	day := "day_" + time.Format("02")
	file := "master_scoreboard.json"
	return host + "/" + main + "/" + year + "/" + month + "/" + day + "/" + file
}
