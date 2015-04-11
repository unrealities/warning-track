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
	Outs         int
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

func BaseOut(outs int, br1 bool, br2 bool, br3 bool) int {
	baseOuts := BaseOuts()
	bo := 0

	for key, value := range baseOuts {
		if value.Outs == outs &&
			value.Base_Runners.First == br1 &&
			value.Base_Runners.Second == br2 &&
			value.Base_Runners.Third == br3 {
			bo = key
		}
	}

	return bo
}
