package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func GameJSON(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	liveGames := []game{}

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Game")
	t := q.Run(c)
	for {
		var g game
		_, err := t.Next(&g)
		if err == datastore.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		liveGames = append(liveGames, g)
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(liveGames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func SetGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// If before 12pm UTC (8am EST). Display the results from the day before
	gameTime := time.Now().UTC()
	if gameTime.Hour() < 12 {
		gameTime = time.Now().UTC().Add(-12 * time.Hour)
	}

	c := appengine.NewContext(r)

	// Delete existing games
	deleteQuery := datastore.NewQuery("Game").KeysOnly()
	deleteKeys, deleteQueryErr := deleteQuery.GetAll(c, nil)
	if deleteQueryErr != nil {
		http.Error(w, deleteQueryErr.Error(), http.StatusInternalServerError)
		return
	}
	deleteErr := datastore.DeleteMulti(c, deleteKeys)
	if deleteErr != nil {
		http.Error(w, deleteErr.Error(), http.StatusInternalServerError)
		return
	}

	// Prevent going over quota
	time.Sleep(2000 * time.Millisecond)

	liveGames := []game{}
	teams := Teams()

	msb := MasterScoreboard(gameTime, r)
	for _, g := range msb.Data.Games.Game {
		outs, _ := strconv.Atoi(g.GameStatus.Outs)
		base_runners, _ := strconv.Atoi(g.RunnersOnBase.Status)

		home_team_runs, _ := strconv.Atoi(g.LineScore.Runs.Home)
		away_team_runs, _ := strconv.Atoi(g.LineScore.Runs.Away)
		run_diff := home_team_runs - away_team_runs

		inning, _ := strconv.Atoi(g.GameStatus.Inning)
		top := false
		if g.GameStatus.TopInning == "Y" {
			top = true
		}

		if run_diff > 4 {
			run_diff = 4
		}
		if run_diff < -4 {
			run_diff = -4
		}

		if outs > 2 && top == true {
			outs = 0
			top = false
		}
		if outs > 2 && top == false {
			outs = 0
			top = true
			inning++
		}
		if inning > 9 {
			inning = 9
		}

		bo := (outs + 1) * base_runners
		gs := GameState(inning, top, run_diff)
		li := 0.0
		if g.GameStatus.Status == "In Progress" {
			li = LeverageIndex(bo, gs)
		}
		g.Li = li

		for _, t := range teams {
			if t.Abbr == g.HomeTeamAbbr {
				g.HomeTeam = t
			} else if t.Abbr == g.AwayTeamAbbr {
				g.AwayTeam = t
			}
		}

		liveGames = append(liveGames, g)
	}

	keys := make([]*datastore.Key, len(liveGames))
	for key := range keys {
		keys[key] = datastore.NewKey(c, "Game", "", 0, nil)
	}

	_, err := datastore.PutMulti(c, keys, liveGames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(liveGames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func Routes() http.Handler {
	router := httprouter.New()

	router.GET("/games", GameJSON)
	router.GET("/fetchGames", SetGames)
	router.NotFound = http.FileServer(http.Dir("static/")).ServeHTTP

	return router
}
