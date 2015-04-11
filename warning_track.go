package main

/*
MLB Leverage Index

The values are from: http://www.insidethebook.com/li.shtml

*/

import (
	"fmt"
	"strconv"
)

func main() {
	gameEvents := GameEvents()
	run_diff := 0
	inning := 1
	top := true
	bo := 0
	gs := 0
	li := LeverageIndex(bo, gs)

	for _, val := range gameEvents.Data.Game.Inning {
		for _, t := range val.Top.Atbat {
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
				top = true

				gs = GameState(inning, top, run_diff)

				li = LeverageIndex(bo, gs)
			}
		}

		li = LeverageIndex(bo, gs)

		for _, b := range val.Bottom.Atbat {
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

	fmt.Println(strconv.FormatFloat(li, 'f', -1, 64))
}
