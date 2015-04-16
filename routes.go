package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO: This needs to be smarter for timezones and past midnight
	// DEBUG: time.Date(2015, time.April, 15, 23, 0, 0, 0, time.UTC)
	gameTime := time.Now()

	liveGames := []gameInfo{}

	grids := Grid(gameTime, r)
	for _, game := range grids.Data.Games.Game {
		if game.Status != "In Progress" {
			continue
		}

		gameEvents := GameEvents(gameTime, game.Id, r)
		outs := 0
		br1, br2, br3 := false, false, false
		run_diff := 0
		inning := 1
		top := true
		bo := 0
		gs := 0
		li := LeverageIndex(bo, gs)

		for _, val := range gameEvents.Data.Game.Innings {
			for _, t := range val.Top.AtBats {
				outs, _ = strconv.Atoi(t.O)
				br1, br2, br3 = false, false, false

				if t.B1 > "" {
					br1 = true
				}
				if t.B2 > "" {
					br2 = true
				}
				if t.B3 > "" {
					br3 = true
				}

				home, _ := strconv.Atoi(t.Home_Team_Runs)
				away, _ := strconv.Atoi(t.Away_Team_Runs)
				run_diff = home - away

				top = true
				inning, _ = strconv.Atoi(val.Num)

				li = CalcLeverageIndex(outs, br1, br2, br3, inning, top, run_diff)
			}

			li = CalcLeverageIndex(outs, br1, br2, br3, inning, top, run_diff)

			for _, b := range val.Bottom.AtBats {
				outs, _ = strconv.Atoi(b.O)
				br1, br2, br3 = false, false, false

				if b.B1 > "" {
					br1 = true
				}
				if b.B2 > "" {
					br2 = true
				}
				if b.B3 > "" {
					br3 = true
				}

				home, _ := strconv.Atoi(b.Home_Team_Runs)
				away, _ := strconv.Atoi(b.Away_Team_Runs)
				run_diff = home - away

				top = false
				inning, _ = strconv.Atoi(val.Num)

				li = CalcLeverageIndex(outs, br1, br2, br3, inning, top, run_diff)
			}
			li = CalcLeverageIndex(outs, br1, br2, br3, inning, top, run_diff)
		}
		newGame := gameInfo{game.Id, game.Status, li}
		liveGames = append(liveGames, newGame)
	}

	sort.Sort(ByLi(liveGames))
	for _, g := range liveGames {
		fmt.Fprintln(w, g)
	}
}

func Routes() http.Handler {
	router := httprouter.New()

	router.GET("/", Index)
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	return router
}
