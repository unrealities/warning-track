package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
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

func Grid(time time.Time, r *http.Request) grid {
	grids := grid{}

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(GridURL(time))
	if err != nil {
		fmt.Println("Error accessing GridURL: " + err.Error())
		panic(err)
	}
	defer resp.Body.Close()

	gridsData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading gridsData: " + err.Error())
	}

	err = json.Unmarshal(gridsData, &grids)
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
