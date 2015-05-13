package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
)

func GameJSON(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)

	liveGames := []game{}
	liveStatuses := []status{}

	_, get_cache_err := memcache.JSON.Get(c, "Game", &liveGames)
	if get_cache_err != nil && get_cache_err != memcache.ErrCacheMiss {
		http.Error(w, get_cache_err.Error(), http.StatusInternalServerError)
	}
	if get_cache_err == nil {
		//success
	} else {
		q := datastore.NewQuery("Game")

		_, Err := q.GetAll(c, &liveGames)
		if Err != nil {
			http.Error(w, Err.Error(), http.StatusInternalServerError)
			return
		}

		item := &memcache.Item{
			Key:    "Game",
			Object: liveGames,
		}
		setErr := memcache.JSON.Set(c, item)
		if setErr != nil {
			http.Error(w, setErr.Error(), http.StatusInternalServerError)
		}
	}

	warningTrackGames := make([]wtGame, len(liveGames))

	_, get_cache_err = memcache.JSON.Get(c, "Status", &liveStatuses)
	if get_cache_err != nil && get_cache_err != memcache.ErrCacheMiss {
		http.Error(w, get_cache_err.Error(), http.StatusInternalServerError)
	}
	if get_cache_err == nil {
		//success
	} else {
		q := datastore.NewQuery("Status")

		_, Err := q.GetAll(c, &liveStatuses)
		if Err != nil {
			http.Error(w, Err.Error(), http.StatusInternalServerError)
			return
		}

		item := &memcache.Item{
			Key:    "Status",
			Object: liveStatuses,
		}
		setErr := memcache.JSON.Set(c, item)
		if setErr != nil {
			http.Error(w, setErr.Error(), http.StatusInternalServerError)
		}
	}

	for k, lg := range liveGames {
		warningTrackGames[k].Id = lg.Id
		warningTrackGames[k].Teams = lg.Teams
		warningTrackGames[k].Links = lg.Links

		t, _ := time.Parse("2006/01/02 3:04PM MST", lg.DateTime)
		warningTrackGames[k].DateTime = t.UTC().Format(time.RFC3339)

		for _, ls := range liveStatuses {
			if ls.GameId == lg.Id {
				warningTrackGames[k].Status = ls
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := JSONMarshal(warningTrackGames, true)
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
		g.DateTime = m.TimeDate + m.AmPm + " EST"
		g.Links.MlbTv = mlbApiMlbTvLinkToUrl(m.Links.MlbTv)

		games = append(games, g)
	}

	// store games
	item := &memcache.Item{
		Key:    "Game",
		Object: games,
	}
	setErr := memcache.JSON.Set(c, item)
	if setErr != nil {
		http.Error(w, setErr.Error(), http.StatusInternalServerError)
	}

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

	ls := []status{}

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
		if g.GameStatus.Status == "In Progress" || g.GameStatus.Status == "Manager Challenge" {
			li = LeverageIndex(bo, gs)
		}

		//convert from mlbApiGame to status
		s := status{}
		s.GameId, _ = strconv.Atoi(g.GamePk)
		s.State = gameStateToInt(g.GameStatus.Status)
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
			a := alert{}
			teams := Teams()

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
			a.Link = mlbApiMlbTvLinkToUrl(g.Links.MlbTv)
			a.Batter = g.Batter.Last

			alertMessage := AlertMessage(a)
			Tweet(alertMessage, w, r)
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
		if g.GameStatus.Status == "In Progress" || g.GameStatus.Status == "Manager Challenge" {
			li = LeverageIndex(bo, gs)
		}

		//convert from mlbApiGame to status
		s := status{}
		s.GameId, _ = strconv.Atoi(g.GamePk)
		s.State = gameStateToInt(g.GameStatus.Status)
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

func DeleteGamesStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)

	//delete games
	g := []game{}
	gq := datastore.NewQuery("Game").KeysOnly()

	gk, Err := gq.GetAll(c, &g)
	if Err != nil {
		http.Error(w, Err.Error(), http.StatusInternalServerError)
		return
	}

	datastore.DeleteMulti(c, gk)
	memcache.Delete(c, "Game")

	//delete statuses
	s := []status{}
	sq := datastore.NewQuery("Status").KeysOnly()

	sk, Err := sq.GetAll(c, &s)
	if Err != nil {
		http.Error(w, Err.Error(), http.StatusInternalServerError)
		return
	}

	datastore.DeleteMulti(c, sk)
	memcache.Delete(c, "Status")

	//reload current games and statuses
	SetGames(w, r, nil)
	SetAllStatuses(w, r, nil)
}

func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
}

func Routes() http.Handler {
	router := httprouter.New()

	router.GET("/games", GameJSON)
	router.GET("/fetchGames", SetGames)
	router.GET("/fetchStatuses", SetStatuses)
	router.GET("/fetchAllStatuses", SetAllStatuses)
	router.GET("/setTwitterCredentials", SetTwitterCredentials)
	router.GET("/deleteGamesStatuses", DeleteGamesStatuses)
	router.NotFound = http.FileServer(http.Dir("static/")).ServeHTTP

	return router
}
