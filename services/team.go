package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/unrealities/warning-track/models"
)

func Teams() []models.Team {
	teams := []models.Team{}

	teamsFile, err := os.Open("team.json")
	if err != nil {
		fmt.Println("Error opening teamsFile: " + err.Error())
	}

	jsonParser := json.NewDecoder(teamsFile)
	if err = jsonParser.Decode(&teams); err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return teams
}
