package main

type game struct {
	HomeTeamAbbr  string     `json:"home_name_abbrev"`
	AwayTeamAbbr  string     `json:"away_name_abbrev"`
	GameStatus    gameStatus `json:"status"`
	RunnersOnBase struct {
		// 0:none; 1:1b; 2:2b; 3:3b; 4:1b,2b; 5:1b,3b; 6:2b,3b; 7:1b,2b,3b
		Status string `json:"status"`
	} `json:"runners_on_base"`
	LineScore lineScore `json:"linescore"`
	Li        float64   `json:"leverage_index"`
	HomeTeam  team      `json:"home_team"`
	AwayTeam  team      `json:"away_team"`
	Links     links     `json:"links"`
}

type gameStatus struct {
	TopInning string `json:"top_inning"`
	Strikes   string `json:"s"`
	Balls     string `json:"b"`
	Status    string `json:"status"`
	Outs      string `json:"o"`
	Inning    string `json:"inning"`
}

type lineScore struct {
	Runs struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"r"`
}

type links struct {
	MlbTv string `json:"mlbtv"`
}
