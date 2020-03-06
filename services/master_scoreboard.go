package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/unrealities/warning-track/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func MasterScoreboard(time time.Time, r *http.Request) models.Scoreboard {
	scoreboards := models.Scoreboard{}

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
	host := "http://statsapi.mlb.com"
	path := "/api/v1/schedule"
	query := "?sportId=1&hydrate=game(content(summary,media(epg))),linescore(runners),flags,team&date="
	month := time.Format("01")
	day := time.Format("02")
	year := time.Format("2006")
	return host + path + query + month + "/" + day + "/" + year
}
