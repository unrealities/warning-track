package models

type Scoreboard struct {
	Data struct {
		Games struct {
			Game []mlbApiGame `json:"game"`
		}
	}
}
