package main

type game struct {
	Id       int    `json:"id"`
	Teams    teams  `json:"teams"`
	DateTime string `json:"date_time"`
	Links    links  `json:"links"`
}

type status struct {
	GameId          int     `json:"game_id"`
	State           string  `json:"state"`
	Score           score   `json:"score"`
	BaseRunnerState int     `json:"base_runner_state"` // 0:none; 1:1b; 2:2b; 3:3b; 4:1b,2b; 5:1b,3b; 6:2b,3b; 7:1b,2b,3b
	Inning          int     `json:"inning"`
	HalfInning      string  `json:"half_inning"` // "top" or "bottom"
	Count           count   `json:"count"`
	Outs            int     `json:"outs"`
	Li              float64 `json:"leverage_index"`
}

type teams struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type links struct {
	MlbTv string `json:"mlb_tv"`
}

type score struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type count struct {
	Balls   int `json:"balls"`
	Strikes int `json:"strikes"`
}
