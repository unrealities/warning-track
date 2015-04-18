package main

import "fmt"

func (g game) String() string {
	return fmt.Sprintf("%s: %-1.1f", g.GameDataDirectory, g.Li)
}

type game struct {
	HomeTeamCity      string     `json:"home_team_city"`
	HomeTeamName      string     `json:"home_team_name"`
	AwayTeamCity      string     `json:"away_team_city"`
	AwayTeamName      string     `json:"away_team_name"`
	Venue             string     `json:"venue"`
	GameDataDirectory string     `json:"game_data_directory"`
	GameStatus        gameStatus `json:"status"`
	RunnersOnBase     struct {
		// 0:none; 1:1b; 2:2b; 3:3b; 4:1b,2b; 5:1b,3b; 6:2b,3b; 7:1b,2b,3b
		Status string `json:"status"`
	} `json:"runners_on_base"`
	LineScore lineScore `json:"linescore"`
	Pitcher   pitcher   `json:"pitcher"`
	Batter    batter    `json:"batter"`
	Li        float64   `json:"leverage_index"`
}

type gameStatus struct {
	TopInning   string `json:"top_inning"`
	Strikes     string `json:"s"`
	Balls       string `json:"b"`
	Status      string `json:"status"`
	Outs        string `json:"o"`
	Inning      string `json:"inning"`
	InningState string `json:"inning_state"`
}

type lineScore struct {
	Runs struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"r"`
	Hits struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"h"`
	Errors struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"e"`
}

type pitcher struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

type batter struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}
