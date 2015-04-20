package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type team struct {
	id     int64
	mlb_id int64
	abbr   string
}

func Teams() []team {
	teams := []team{}

	teamsFile, err := os.Open("teams.json")
	if err != nil {
		fmt.Println("Error opening teamsFile: " + err.Error())
	}

	jsonParser := json.NewDecoder(teamsFile)
	if err = jsonParser.Decode(&teams); err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return teams
}
