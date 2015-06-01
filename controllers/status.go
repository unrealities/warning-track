package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/unrealities/warning-track/models"
	"github.com/unrealities/warning-track/services"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
)

func SetStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)
	// If before 12pm UTC (8am EST). Display the results from the day before
	gameTime := time.Now().UTC()
	if gameTime.Hour() < 12 {
		gameTime = time.Now().UTC().Add(-12 * time.Hour)
	}

	ls := []models.Status{}

	// Fetch "In Progress" games
	q := datastore.NewQuery("Status").
		Filter("State >", 20).
		Project("GameId")

	_, Err := q.GetAll(c, &ls)
	if Err != nil {
		http.Error(w, Err.Error(), http.StatusInternalServerError)
		return
	}

	if len(ls) == 0 {
		w.Write([]byte("[]"))
		return
	}

	// store game ids for updating
	liveGameIds := []int{}
	for _, liveStatus := range ls {
		liveGameIds = append(liveGameIds, liveStatus.GameId)
	}

	statuses := []models.Status{}
	msb := services.MasterScoreboard(gameTime, r)
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

		init_outs, _ := strconv.Atoi(g.GameStatus.Outs)
		outs := init_outs
		base_runners, _ := strconv.Atoi(g.RunnersOnBase.Status)

		home_team_runs, _ := strconv.Atoi(g.LineScore.Runs.Home)
		away_team_runs, _ := strconv.Atoi(g.LineScore.Runs.Away)
		init_run_diff := home_team_runs - away_team_runs
		run_diff := init_run_diff

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
		gs := services.GameState(inning, top, run_diff)
		li := 0.0
		if g.GameStatus.Status == "In Progress" || g.GameStatus.Status == "Manager Challenge" {
			li = services.LeverageIndex(bo, gs)
		}
		if init_run_diff > 4 || init_run_diff < -4 || init_outs > 2 {
			li = 0.0
		}

		//convert from mlbApiGame to status
		s := models.Status{}
		s.GameId, _ = strconv.Atoi(g.GamePk)
		s.State = services.GameStateToInt(g.GameStatus.Status)
		s.Score.Home = home_team_runs
		s.Score.Away = away_team_runs
		s.BaseRunnerState = base_runners
		s.Inning, _ = strconv.Atoi(g.GameStatus.Inning)
		s.HalfInning = "Bot"
		if g.GameStatus.TopInning == "Y" {
			s.HalfInning = "Top"
		}
		s.Count.Balls, _ = strconv.Atoi(g.GameStatus.Balls)
		s.Count.Strikes, _ = strconv.Atoi(g.GameStatus.Strikes)
		s.Outs, _ = strconv.Atoi(g.GameStatus.Outs)
		s.Li = li

		if s.Li >= 3 {
			a := models.Alert{}
			teams := services.Teams()

			for _, t := range teams {
				if t.Abbr == g.HomeTeamAbbr {
					a.Teams.Home = t.Id
				} else if t.Abbr == g.AwayTeamAbbr {
					a.Teams.Away = t.Id
				}
			}
			a.Score.Home = s.Score.Home
			a.Score.Away = s.Score.Away
			a.Inning = s.Inning
			a.HalfInning = s.HalfInning
			a.Outs = s.Outs
			a.BaseRunnerState = s.BaseRunnerState
			a.Li = s.Li
			a.Link = services.MlbApiMlbTvLinkToUrl(g.Links.MlbTv)
			a.Batter = g.Batter.Last

			alertMessage := services.AlertMessage(a)
			services.Tweet(alertMessage, w, r)
		}

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

	memcache.Delete(c, "Status")

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

	statuses := []models.Status{}
	msb := services.MasterScoreboard(gameTime, r)
	for _, g := range msb.Data.Games.Game {
		init_outs, _ := strconv.Atoi(g.GameStatus.Outs)
		outs := init_outs
		base_runners, _ := strconv.Atoi(g.RunnersOnBase.Status)

		home_team_runs, _ := strconv.Atoi(g.LineScore.Runs.Home)
		away_team_runs, _ := strconv.Atoi(g.LineScore.Runs.Away)
		init_run_diff := home_team_runs - away_team_runs
		run_diff := init_run_diff

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
		gs := services.GameState(inning, top, run_diff)
		li := 0.0
		if g.GameStatus.Status == "In Progress" || g.GameStatus.Status == "Manager Challenge" {
			li = services.LeverageIndex(bo, gs)
		}
		if init_run_diff > 4 || init_run_diff < -4 || init_outs > 2 {
			li = 0.0
		}

		//convert from mlbApiGame to status
		s := models.Status{}
		s.GameId, _ = strconv.Atoi(g.GamePk)
		s.State = services.GameStateToInt(g.GameStatus.Status)
		s.Score.Home = home_team_runs
		s.Score.Away = away_team_runs
		s.BaseRunnerState = base_runners
		s.Inning, _ = strconv.Atoi(g.GameStatus.Inning)
		s.HalfInning = "Bot"
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
	item := &memcache.Item{
		Key:    "Status",
		Object: statuses,
	}
	setErr := memcache.JSON.Set(c, item)
	if setErr != nil {
		http.Error(w, setErr.Error(), http.StatusInternalServerError)
	}

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
