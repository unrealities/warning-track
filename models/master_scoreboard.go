package models

import "time"

type Scoreboard struct {
	Dates []struct {
		Date  string `json:"date"`
		Games []struct {
			GamePk   int       `json:"gamePk"`
			GameType string    `json:"gameType"`
			Season   string    `json:"season"`
			GameDate time.Time `json:"gameDate"`
			Teams    struct {
				Away struct {
					Team struct {
						ID           int64  `json:"id"`
						Abbreviation string `json:"abbreviation"`
					} `json:"team"`
				} `json:"away"`
				Home struct {
					Team struct {
						ID           int64  `json:"id"`
						Abbreviation string `json:"abbreviation"`
					} `json:"team"`
				} `json:"home"`
			} `json:"teams"`
			Linescore struct {
				CurrentInning        int    `json:"currentInning"`
				CurrentInningOrdinal string `json:"currentInningOrdinal"`
				InningState          string `json:"inningState"`
				InningHalf           string `json:"inningHalf"`
				IsTopInning          bool   `json:"isTopInning"`
				ScheduledInnings     int    `json:"scheduledInnings"`
				Innings              []struct {
					Num        int    `json:"num"`
					OrdinalNum string `json:"ordinalNum"`
					Home       struct {
						Runs       int `json:"runs"`
						Hits       int `json:"hits"`
						Errors     int `json:"errors"`
						LeftOnBase int `json:"leftOnBase"`
					} `json:"home"`
					Away struct {
						Runs       int `json:"runs"`
						Hits       int `json:"hits"`
						Errors     int `json:"errors"`
						LeftOnBase int `json:"leftOnBase"`
					} `json:"away"`
				} `json:"innings"`
				Teams struct {
					Home struct {
						Runs       int `json:"runs"`
						Hits       int `json:"hits"`
						Errors     int `json:"errors"`
						LeftOnBase int `json:"leftOnBase"`
					} `json:"home"`
					Away struct {
						Runs       int `json:"runs"`
						Hits       int `json:"hits"`
						Errors     int `json:"errors"`
						LeftOnBase int `json:"leftOnBase"`
					} `json:"away"`
				} `json:"teams"`
				Defense struct {
					Batter struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"batter"`
					OnDeck struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"onDeck"`
					InHole struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"inHole"`
				} `json:"defense"`
				Offense Offense `json:"offense"`
				Balls   int     `json:"balls"`
				Strikes int     `json:"strikes"`
				Outs    int     `json:"outs"`
			} `json:"linescore,omitempty"`
			Content struct {
				Media struct {
					Epg []struct {
						Title string `json:"title"`
						Items []struct {
							ID               int    `json:"id"`
							ContentID        string `json:"contentId"`
							CallLetters      string `json:"callLetters"`
							FoxAuthRequired  bool   `json:"foxAuthRequired"`
							TbsAuthRequired  bool   `json:"tbsAuthRequired"`
							EspnAuthRequired bool   `json:"espnAuthRequired"`
							Fs1AuthRequired  bool   `json:"fs1AuthRequired"`
							MlbnAuthRequired bool   `json:"mlbnAuthRequired"`
							FreeGame         bool   `json:"freeGame"`
						} `json:"items"`
						FreeGame     bool `json:"freeGame"`
						EnhancedGame bool `json:"enhancedGame"`
					} `json:"epg"`
				} `json:"media"`
			} `json:"content"`
			Status struct {
				AbstractGameState string `json:"abstractGameState"`
				CodedGameState    string `json:"codedGameState"`
				DetailedState     string `json:"detailedState"`
				StatusCode        string `json:"statusCode"`
				AbstractGameCode  string `json:"abstractGameCode"`
			} `json:"status"`
		} `json:"games"`
	} `json:"dates"`
}

type Offense struct {
	First struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"first,omitempty"`
	Second struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"second,omitempty"`
	Third struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"third,omitempty"`
}
