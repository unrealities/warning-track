package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/unrealities/warning-track/models"
)

func BaseOuts() []models.BaseOut {
	baseOuts := []models.BaseOut{}

	baseOutsFile, err := os.Open("../models/base_out.json")
	if err != nil {
		fmt.Println("Error opening baseOutsFile: " + err.Error())
	}

	jsonParser := json.NewDecoder(baseOutsFile)
	if err = jsonParser.Decode(&baseOuts); err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return baseOuts
}
