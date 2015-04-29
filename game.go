package main

type game struct {
	HomeTeamAbbr      string     `json:"home_name_abbrev"`
	AwayTeamAbbr      string     `json:"away_name_abbrev"`
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
	Pbp       pbp       `json:"pbp"`
	Li        float64   `json:"leverage_index"`
	HomeTeam  team      `json:"home_team"`
	AwayTeam  team      `json:"away_team"`
	Links     links     `json:"links"`
}

type gameStatus struct {
	TopInning   string `json:"top_inning"`
	Strikes     string `json:"s"`
	Balls       string `json:"b"`
	Status      string `json:"status"`
	Outs        string `json:"o"`
	Inning      string `json:"inning"`
	InningState string `json:"inning_state"`
	Note        string `json:"note"`
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

type pbp struct {
	Last string `json:"last"`
}

type links struct {
	AwayAudio string `json:"away_audio"`
	HomeAudio string `json:"home_audio"`
	WrapUp    string `json:"wrapup"`
	Preview   string `json:"preview"`
	TVStation string `json:"tv_station"`
	MlbTv     string `json:"mlbtv"`
}
