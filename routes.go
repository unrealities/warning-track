package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func GameJSON(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// If before 12pm UTC (8am EST). Display the results from the day before
	// DEBUG: time.Date(2015, time.April, 15, 23, 0, 0, 0, time.UTC)
	gameTime := time.Now().UTC()
	if gameTime.Hour() < 12 {
		gameTime = time.Now().UTC().Add(-12 * time.Hour)
	}

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
		li := LeverageIndex(bo, gs)
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
	router.NotFound = http.FileServer(http.Dir("static/")).ServeHTTP

	return router
}
