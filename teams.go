package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type team struct {
	Id      int    `json:"id"`
	MlbId   int64  `json:"mlb_id"`
	Abbr    string `json:"abbr"`
	Hashtag string `json:"hashtag"`
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
