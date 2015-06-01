package models

type Alert struct {
	CreatedAt       string  `json:"created_at"`
	Teams           teams   `json:"teams"`
	Score           score   `json:"score"`
	Inning          int     `json:"inning"`
	HalfInning      string  `json:"half_inning"`
	Outs            int     `json:"outs"`
	BaseRunnerState int     `json:"base_runner_state"`
	Li              float64 `json:"leverage_index"`
	Link            string  `json:"link"`
	Batter          string  `json:"batter"`
}
