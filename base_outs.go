package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type baseRunners struct {
	First  bool
	Second bool
	Third  bool
}

type baseOut struct {
	Outs         int8
	Base_Runners baseRunners
}

func BaseOuts() []baseOut {
	baseOuts := []baseOut{}

	baseOutsFile, err := os.Open("base_outs.json")
	if err != nil {
		fmt.Println("Error opening baseOutsFile: " + err.Error())
	}

	jsonParser := json.NewDecoder(baseOutsFile)
	if err = jsonParser.Decode(&baseOuts); err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return baseOuts
}
