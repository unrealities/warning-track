package main

/*
MLB Leverage Index

The values are from: http://www.insidethebook.com/li.shtml
*/

import "fmt"

func main() {
	baseOuts := BaseOuts()
	fmt.Println("Base-outs")
	fmt.Println(baseOuts)

	gameStates := GameStates()
	fmt.Println("Game States")
	fmt.Println(gameStates)

	gameEvents := GameEvents()
	fmt.Println(gameEvents.Subject)

	for _, val := range gameEvents.Data.Game.Inning {
		fmt.Println("Top of " + val.Num)
		for _, t := range val.Top.Atbat {
			fmt.Printf("Score: ")
			fmt.Printf(t.Away_Team_Runs + " - ")
			fmt.Printf(t.Home_Team_Runs)
			fmt.Println()
			fmt.Printf("Outs: ")
			fmt.Printf(t.O + " ")
			fmt.Println()
			fmt.Printf("B1: " + t.B1 + " ")
			fmt.Printf("B2: " + t.B2 + " ")
			fmt.Printf("B3: " + t.B3 + " ")
			fmt.Println()
		}
		fmt.Println()

		fmt.Println("Bottom of " + val.Num)
		for _, b := range val.Bottom.Atbat {
			fmt.Printf("Score: ")
			fmt.Printf(b.Away_Team_Runs + " - ")
			fmt.Printf(b.Home_Team_Runs)
			fmt.Println()
			fmt.Printf("Outs: ")
			fmt.Printf(b.O + " ")
			fmt.Println()
			fmt.Printf("B1: " + b.B1 + " ")
			fmt.Printf("B2: " + b.B2 + " ")
			fmt.Printf("B3: " + b.B3 + " ")
			fmt.Println()
		}
		fmt.Println()
	}
}
