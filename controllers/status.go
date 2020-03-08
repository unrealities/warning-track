package controllers

import (
	"encoding/json"
	"net/http"
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
	for _, g := range msb.Dates[0].Games {
		// only run for live games
		update := false
		for _, l := range liveGameIds {
			if l == g.GamePk {
				update = true
				break
			}
		}
		if update == false {
			continue
		}

		initOuts := g.Linescore.Outs
		outs := initOuts
		baseRunnerState := offenseToBaseRunnerState(g.Linescore.Offense)

		htRuns := g.Linescore.Teams.Home.Runs
		atRuns := g.Linescore.Teams.Away.Runs
		runDiff := htRuns - atRuns

		inning := g.Linescore.CurrentInning
		top := false
		if g.Linescore.IsTopInning {
			top = true
		}

		if runDiff > 4 {
			runDiff = 4
		}
		if runDiff < -4 {
			runDiff = -4
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

		bo := (outs + 1) * baseRunnerState
		gs := services.GameState(inning, top, runDiff)
		li := 0.0
		if g.Status.DetailedState == "In Progress" || g.Status.DetailedState == "Manager Challenge" {
			li = services.LeverageIndex(bo, gs)
		}
		if runDiff >= 4 || runDiff <= -4 || initOuts > 2 {
			li = 0.0
		}

		//convert from mlbApiGame to status
		s := models.Status{}
		s.GameId = g.GamePk
		s.State = services.GameStateToInt(g.Status.DetailedState)
		s.Score.Home = htRuns
		s.Score.Away = atRuns
		s.BaseRunnerState = baseRunnerState
		s.Inning = g.Linescore.CurrentInning
		s.HalfInning = "Bot"
		if top {
			s.HalfInning = "Top"
		}
		s.Count.Balls = g.Linescore.Balls
		s.Count.Strikes = g.Linescore.Strikes
		s.Outs = g.Linescore.Outs
		s.Li = li

		if s.Li >= 3 {
			a := models.Alert{}
			teams := services.Teams()

			for _, t := range teams {
				if t.MlbId == g.Teams.Home.Team.ID {
					a.Teams.Home = t.Id
				} else if t.MlbId == g.Teams.Away.Team.ID {
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
			a.Batter = g.Linescore.Defense.Batter.FullName
			a.Link = services.MlbApiMlbTvLinkToUrl(g.GamePk)

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
	for _, g := range msb.Dates[0].Games {
		initOuts := g.Linescore.Outs
		outs := initOuts
		baseRunnerState := offenseToBaseRunnerState(g.Linescore.Offense)

		htRuns := g.Linescore.Teams.Home.Runs
		atRuns := g.Linescore.Teams.Away.Runs
		runDiff := htRuns - atRuns

		inning := g.Linescore.CurrentInning
		top := false
		if g.Linescore.IsTopInning {
			top = true
		}

		if runDiff > 4 {
			runDiff = 4
		}
		if runDiff < -4 {
			runDiff = -4
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

		bo := (outs + 1) * baseRunnerState
		gs := services.GameState(inning, top, runDiff)
		li := 0.0
		if g.Status.DetailedState == "In Progress" || g.Status.DetailedState == "Manager Challenge" {
			li = services.LeverageIndex(bo, gs)
		}
		if runDiff >= 4 || runDiff <= -4 || initOuts > 2 {
			li = 0.0
		}

		//convert from mlbApiGame to status
		s := models.Status{}
		s.GameId = g.GamePk
		s.State = services.GameStateToInt(g.Status.DetailedState)
		s.Score.Home = htRuns
		s.Score.Away = atRuns
		s.BaseRunnerState = baseRunnerState
		s.Inning = g.Linescore.CurrentInning
		s.HalfInning = "Bot"
		if top {
			s.HalfInning = "Top"
		}
		s.Count.Balls = g.Linescore.Balls
		s.Count.Strikes = g.Linescore.Strikes
		s.Outs = g.Linescore.Outs
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

// 0:none; 1:1b; 2:2b; 3:3b; 4:1b,2b; 5:1b,3b; 6:2b,3b; 7:1b,2b,3b
func offenseToBaseRunnerState(o models.Offense) int {
	switch {
	case o.First.ID == 0 && o.Second.ID == 0 && o.Third.ID == 0:
		return 0
	case o.First.ID > 0 && o.Second.ID == 0 && o.Third.ID == 0:
		return 1
	case o.First.ID == 0 && o.Second.ID > 0 && o.Third.ID == 0:
		return 2
	case o.First.ID == 0 && o.Second.ID == 0 && o.Third.ID > 0:
		return 3
	case o.First.ID > 0 && o.Second.ID > 0 && o.Third.ID == 0:
		return 4
	case o.First.ID > 0 && o.Second.ID == 0 && o.Third.ID > 0:
		return 5
	case o.First.ID == 0 && o.Second.ID > 0 && o.Third.ID > 0:
		return 6
	case o.First.ID > 0 && o.Second.ID > 0 && o.Third.ID > 0:
		return 7
	}
	return 0
}
