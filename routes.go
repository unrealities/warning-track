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
	c := appengine.NewContext(r)

	liveGames := []game{}
	liveStatuses := []status{}

	gq := datastore.NewQuery("Game")
	t := gq.Run(c)
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
	warningTrackGames := make([]wtGame, len(liveGames))

	sq := datastore.NewQuery("Status")
	t = sq.Run(c)
	for {
		var s status
		_, err := t.Next(&s)
		if err == datastore.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		liveStatuses = append(liveStatuses, s)
	}

	for k, lg := range liveGames {
		warningTrackGames[k].Id = lg.Id
		warningTrackGames[k].Teams = lg.Teams
		warningTrackGames[k].Links = lg.Links
		for _, ls := range liveStatuses {
			if ls.GameId == lg.Id {
				warningTrackGames[k].Status = ls
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(warningTrackGames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func SetGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)
	// If before 12pm UTC (8am EST). Display the results from the day before
	gameTime := time.Now().UTC()
	if gameTime.Hour() < 12 {
		gameTime = time.Now().UTC().Add(-12 * time.Hour)
	}

	teams := Teams()
	games := []game{}
	msb := MasterScoreboard(gameTime, r)
	for _, m := range msb.Data.Games.Game {
		g := game{}
		g.Id, _ = strconv.Atoi(m.GamePk)
		for _, t := range teams {
			if t.Abbr == m.HomeTeamAbbr {
				g.Teams.Home = t.Id
			} else if t.Abbr == m.AwayTeamAbbr {
				g.Teams.Away = t.Id
			}
		}
		g.Links.MlbTv = m.Links.MlbTv

		games = append(games, g)
	}

	// store games
	keys := make([]*datastore.Key, len(games))
	for k := range keys {
		keys[k] = datastore.NewKey(c, "Game", "", int64(games[k].Id), nil)
	}

	_, err := datastore.PutMulti(c, keys, games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func SetStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)
	// If before 12pm UTC (8am EST). Display the results from the day before
	gameTime := time.Now().UTC()
	if gameTime.Hour() < 12 {
		gameTime = time.Now().UTC().Add(-12 * time.Hour)
	}

	// Get existing statuses for "In Progress" games
	//
	// To do this in one query for "Delayed" games,
	// change status to an int and do >=
	q := datastore.NewQuery("Status").
		Filter("State =", "In Progress")
	ls := []status{}
	_, Err := q.GetAll(c, &ls)
	if Err != nil {
		http.Error(w, Err.Error(), http.StatusInternalServerError)
		return
	}

	// store game ids for updating
	liveGameIds := []int{}
	for _, liveStatus := range ls {
		liveGameIds = append(liveGameIds, liveStatus.GameId)
	}

	statuses := []status{}
	msb := MasterScoreboard(gameTime, r)
	for _, g := range msb.Data.Games.Game {
		// only run for live games
		update := false
		for _, l := range liveGameIds {
			if strconv.Itoa(l) == g.GamePk {
				update = true
				break
			}
		}
		if update == false {
			continue
		}

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

		//convert from mlbApiGame to status
		s := status{}
		s.GameId, _ = strconv.Atoi(g.GamePk)
		s.State = g.GameStatus.Status
		s.Score.Home = home_team_runs
		s.Score.Away = away_team_runs
		s.BaseRunnerState = base_runners
		s.Inning, _ = strconv.Atoi(g.GameStatus.Inning)
		s.HalfInning = "Bottom"
		if g.GameStatus.TopInning == "Y" {
			s.HalfInning = "Top"
		}
		s.Count.Balls, _ = strconv.Atoi(g.GameStatus.Balls)
		s.Count.Strikes, _ = strconv.Atoi(g.GameStatus.Strikes)
		s.Outs, _ = strconv.Atoi(g.GameStatus.Outs)
		s.Li = li

		statuses = append(statuses, s)
	}

	// store statuses
	keys := make([]*datastore.Key, len(statuses))
	for k := range keys {
		keys[k] = datastore.NewKey(c, "Status", "", int64(statuses[k].GameId), nil)
	}

	_, err := datastore.PutMulti(c, keys, statuses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(statuses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func SetAllStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)
	// If before 12pm UTC (8am EST). Display the results from the day before
	gameTime := time.Now().UTC()
	if gameTime.Hour() < 12 {
		gameTime = time.Now().UTC().Add(-12 * time.Hour)
	}

	// Get all statuses
	q := datastore.NewQuery("Status")
	ls := []status{}
	_, Err := q.GetAll(c, &ls)
	if Err != nil {
		http.Error(w, Err.Error(), http.StatusInternalServerError)
		return
	}

	statuses := []status{}
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

		//convert from mlbApiGame to status
		s := status{}
		s.GameId, _ = strconv.Atoi(g.GamePk)
		s.State = g.GameStatus.Status
		s.Score.Home = home_team_runs
		s.Score.Away = away_team_runs
		s.BaseRunnerState = base_runners
		s.Inning, _ = strconv.Atoi(g.GameStatus.Inning)
		s.HalfInning = "Bottom"
		if g.GameStatus.TopInning == "Y" {
			s.HalfInning = "Top"
		}
		s.Count.Balls, _ = strconv.Atoi(g.GameStatus.Balls)
		s.Count.Strikes, _ = strconv.Atoi(g.GameStatus.Strikes)
		s.Outs, _ = strconv.Atoi(g.GameStatus.Outs)
		s.Li = li

		statuses = append(statuses, s)
	}

	// store statuses
	keys := make([]*datastore.Key, len(statuses))
	for k := range keys {
		keys[k] = datastore.NewKey(c, "Status", "", int64(statuses[k].GameId), nil)
	}

	_, err := datastore.PutMulti(c, keys, statuses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(statuses)
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
	router.GET("/fetchStatuses", SetStatuses)
	router.GET("/fetchAllStatuses", SetAllStatuses)
	router.NotFound = http.FileServer(http.Dir("static/")).ServeHTTP

	return router
}
