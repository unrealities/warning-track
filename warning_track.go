package main

/*
MLB Leverage Index

The values are from: http://www.insidethebook.com/li.shtml

*/

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type gameInfo struct {
	Id     string
	Status string
	Li     float64
}

func (g gameInfo) String() string {
	return fmt.Sprintf("%s: %-1.1f", g.Id, g.Li)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByLi []gameInfo

func (a ByLi) Len() int           { return len(a) }
func (a ByLi) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLi) Less(i, j int) bool { return a[i].Li > a[j].Li }

func main() {
	// This needs to be smarter for timezones and past midnight
	gameTime := time.Now()

	liveGames := []gameInfo{}

	grids := Grid(gameTime)
	for _, game := range grids.Data.Games.Game {
		if game.Status != "In Progress" {
			continue
		}

		gameEvents := GameEvents(gameTime, game.Id)
		run_diff := 0
		inning := 1
		top := true
		bo := 0
		gs := 0
		li := LeverageIndex(bo, gs)

		for _, val := range gameEvents.Data.Game.Innings {
			for _, t := range val.Top.AtBats {
				outs, _ := strconv.Atoi(t.O)
				br1, br2, br3 := false, false, false

				if t.B1 > "" {
					br1 = true
				}
				if t.B2 > "" {
					br2 = true
				}
				if t.B3 > "" {
					br3 = true
				}

				if outs < 3 {
					bo = BaseOut(outs, br1, br2, br3)

					home, _ := strconv.Atoi(t.Home_Team_Runs)
					away, _ := strconv.Atoi(t.Away_Team_Runs)
					run_diff = home - away
					if run_diff > 4 {
						run_diff = 4
					}
					if run_diff < -4 {
						run_diff = -4
					}
					inning, _ = strconv.Atoi(val.Num)
					if inning > 9 {
						inning = 9
					}
					top = true

					gs = GameState(inning, top, run_diff)

					li = LeverageIndex(bo, gs)
				}
			}

			li = LeverageIndex(bo, gs)

			for _, b := range val.Bottom.AtBats {
				outs, _ := strconv.Atoi(b.O)

				br1, br2, br3 := false, false, false

				if b.B1 > "" {
					br1 = true
				}
				if b.B2 > "" {
					br2 = true
				}
				if b.B3 > "" {
					br3 = true
				}

				if outs < 3 {
					bo = BaseOut(outs, br1, br2, br3)

					home, _ := strconv.Atoi(b.Home_Team_Runs)
					away, _ := strconv.Atoi(b.Away_Team_Runs)
					run_diff = home - away
					if run_diff > 4 {
						run_diff = 4
					}
					if run_diff < -4 {
						run_diff = -4
					}
					top = false

					gs = GameState(inning, top, run_diff)

					li = LeverageIndex(bo, gs)
				}
			}
			li = LeverageIndex(bo, gs)
		}
		newGame := gameInfo{game.Id, game.Status, li}
		liveGames = append(liveGames, newGame)
	}

	sort.Sort(ByLi(liveGames))
	for _, g := range liveGames {
		fmt.Println(g)
	}
}
