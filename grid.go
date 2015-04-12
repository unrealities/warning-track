package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type grid struct {
	Data struct {
		Games struct {
			Game []game
		}
	}
}

type game struct {
	Id     string
	Status string
}

func Grid(time time.Time) grid {
	grids := grid{}

	out, err := os.Create("grid.json")
	if err != nil {
		fmt.Println("Error creating file: " + err.Error())
	}
	defer out.Close()

	resp, err := http.Get(GridURL(time))
	if err != nil {
		fmt.Println("Error accessing file: " + err.Error())
	}
	defer resp.Body.Close()

	gridsFile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading gridsFile: " + err.Error())
	}

	err = json.Unmarshal(gridsFile, &grids)
	if err != nil {
		fmt.Println("Error parsing file: " + err.Error())
	}

	return grids
}

func GridURL(time time.Time) string {
	host := "http://gd2.mlb.com"
	main := "components/game/mlb"
	year := "year_" + time.Format("2006")
	month := "month_" + time.Format("01")
	day := "day_" + time.Format("02")
	file := "grid.json"
	return host + "/" + main + "/" + year + "/" + month + "/" + day + "/" + file
}
