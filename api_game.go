package main

import "regexp"

type mlbApiGame struct {
	GamePk        string           `json:"game_pk"`
	HomeTeamAbbr  string           `json:"home_name_abbrev"`
	AwayTeamAbbr  string           `json:"away_name_abbrev"`
	TimeDate      string           `json:"time_date"`
	AmPm          string           `json:"ampm"`
	GameStatus    mlbApiGameStatus `json:"status"` //"Pre-Game", "Postponed", "Final", "Preview", "Delayed", "Game Over", "In Progress", "Warmup?"
	RunnersOnBase struct {
		Status string `json:"status"` // 0:none; 1:1b; 2:2b; 3:3b; 4:1b,2b; 5:1b,3b; 6:2b,3b; 7:1b,2b,3b
	} `json:"runners_on_base"`
	LineScore mlbApiLineScore `json:"linescore"`
	Li        float64         `json:"leverage_index"`
	Links     mlbApiLinks     `json:"links"`
	Batter    struct {
		Last string `json:"last"`
	} `json:"batter"`
}

type mlbApiGameStatus struct {
	TopInning string `json:"top_inning"`
	Strikes   string `json:"s"`
	Balls     string `json:"b"`
	Status    string `json:"status"`
	Outs      string `json:"o"`
	Inning    string `json:"inning"`
}

type mlbApiLineScore struct {
	Runs struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"r"`
}

type mlbApiLinks struct {
	MlbTv string `json:"mlbtv"`
}

func mlbApiMlbTvLinkToUrl(l string) string {
	calEventId := regexp.MustCompile(`[0-9-]+`)
	c := calEventId.FindString(l)
	return "http://mlb.mlb.com/shared/flash/mediaplayer/v4.5/R7/MP4.jsp?calendar_event_id=" + c + "&content_id=&media_id=&view_key=&media_type=video&source=MLB&sponsor=MLB&clickOrigin=Media+Grid&affiliateId=Media+Grid&team=mlb"
}
