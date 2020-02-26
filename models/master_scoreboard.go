package models

import "time"

type Scoreboard struct {
	Copyright            string `json:"copyright"`
	TotalItems           int    `json:"totalItems"`
	TotalEvents          int    `json:"totalEvents"`
	TotalGames           int    `json:"totalGames"`
	TotalGamesInProgress int    `json:"totalGamesInProgress"`
	Dates                []struct {
		Date                 string `json:"date"`
		TotalItems           int    `json:"totalItems"`
		TotalEvents          int    `json:"totalEvents"`
		TotalGames           int    `json:"totalGames"`
		TotalGamesInProgress int    `json:"totalGamesInProgress"`
		Games                []struct {
			GamePk   int       `json:"gamePk"`
			Link     string    `json:"link"`
			GameType string    `json:"gameType"`
			Season   string    `json:"season"`
			GameDate time.Time `json:"gameDate"`
			Status   struct {
				AbstractGameState string `json:"abstractGameState"`
				CodedGameState    string `json:"codedGameState"`
				DetailedState     string `json:"detailedState"`
				StatusCode        string `json:"statusCode"`
				AbstractGameCode  string `json:"abstractGameCode"`
			} `json:"status"`
			Teams struct {
				Away struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Pct    string `json:"pct"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Link   string `json:"link"`
						Season int    `json:"season"`
						Venue  struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"venue"`
						TeamCode        string `json:"teamCode"`
						FileCode        string `json:"fileCode"`
						Abbreviation    string `json:"abbreviation"`
						TeamName        string `json:"teamName"`
						LocationName    string `json:"locationName"`
						FirstYearOfPlay string `json:"firstYearOfPlay"`
						League          struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"league"`
						Division struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"division"`
						Sport struct {
							ID   int    `json:"id"`
							Link string `json:"link"`
							Name string `json:"name"`
						} `json:"sport"`
						ShortName    string `json:"shortName"`
						SpringLeague struct {
							ID           int    `json:"id"`
							Name         string `json:"name"`
							Link         string `json:"link"`
							Abbreviation string `json:"abbreviation"`
						} `json:"springLeague"`
						AllStarStatus string `json:"allStarStatus"`
						Active        bool   `json:"active"`
					} `json:"team"`
					SplitSquad   bool `json:"splitSquad"`
					SeriesNumber int  `json:"seriesNumber"`
					SpringLeague struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
					} `json:"springLeague"`
				} `json:"away"`
				Home struct {
					LeagueRecord struct {
						Wins   int    `json:"wins"`
						Losses int    `json:"losses"`
						Pct    string `json:"pct"`
					} `json:"leagueRecord"`
					Score int `json:"score"`
					Team  struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Link   string `json:"link"`
						Season int    `json:"season"`
						Venue  struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"venue"`
						TeamCode        string `json:"teamCode"`
						FileCode        string `json:"fileCode"`
						Abbreviation    string `json:"abbreviation"`
						TeamName        string `json:"teamName"`
						LocationName    string `json:"locationName"`
						FirstYearOfPlay string `json:"firstYearOfPlay"`
						League          struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"league"`
						Division struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
							Link string `json:"link"`
						} `json:"division"`
						Sport struct {
							ID   int    `json:"id"`
							Link string `json:"link"`
							Name string `json:"name"`
						} `json:"sport"`
						ShortName    string `json:"shortName"`
						SpringLeague struct {
							ID           int    `json:"id"`
							Name         string `json:"name"`
							Link         string `json:"link"`
							Abbreviation string `json:"abbreviation"`
						} `json:"springLeague"`
						AllStarStatus string `json:"allStarStatus"`
						Active        bool   `json:"active"`
					} `json:"team"`
					SplitSquad   bool `json:"splitSquad"`
					SeriesNumber int  `json:"seriesNumber"`
					SpringLeague struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
					} `json:"springLeague"`
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
				Offense struct {
				} `json:"offense"`
				Balls   int `json:"balls"`
				Strikes int `json:"strikes"`
				Outs    int `json:"outs"`
			} `json:"linescore,omitempty"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"venue"`
			Content struct {
				Link      string `json:"link"`
				Editorial struct {
				} `json:"editorial"`
				Media struct {
					EpgAlternate []struct {
						Items []interface{} `json:"items"`
						Title string        `json:"title"`
					} `json:"epgAlternate"`
					FreeGame     bool `json:"freeGame"`
					EnhancedGame bool `json:"enhancedGame"`
				} `json:"media"`
				Highlights struct {
				} `json:"highlights"`
				Summary struct {
					HasPreviewArticle  bool `json:"hasPreviewArticle"`
					HasRecapArticle    bool `json:"hasRecapArticle"`
					HasWrapArticle     bool `json:"hasWrapArticle"`
					HasHighlightsVideo bool `json:"hasHighlightsVideo"`
				} `json:"summary"`
				GameNotes struct {
				} `json:"gameNotes"`
			} `json:"content"`
			GameNumber        int    `json:"gameNumber"`
			PublicFacing      bool   `json:"publicFacing"`
			DoubleHeader      string `json:"doubleHeader"`
			GamedayType       string `json:"gamedayType"`
			Tiebreaker        string `json:"tiebreaker"`
			CalendarEventID   string `json:"calendarEventID"`
			SeasonDisplay     string `json:"seasonDisplay"`
			DayNight          string `json:"dayNight"`
			ScheduledInnings  int    `json:"scheduledInnings"`
			InningBreakLength int    `json:"inningBreakLength"`
			GamesInSeries     int    `json:"gamesInSeries"`
			SeriesGameNumber  int    `json:"seriesGameNumber"`
			SeriesDescription string `json:"seriesDescription"`
			Flags             struct {
				NoHitter    bool `json:"noHitter"`
				PerfectGame bool `json:"perfectGame"`
			} `json:"flags"`
			RecordSource           string `json:"recordSource"`
			IfNecessary            string `json:"ifNecessary"`
			IfNecessaryDescription string `json:"ifNecessaryDescription"`
			Linescore              struct {
				Note                 string `json:"note"`
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
				Offense struct {
					First struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"first"`
				} `json:"offense"`
				Balls   int `json:"balls"`
				Strikes int `json:"strikes"`
				Outs    int `json:"outs"`
			} `json:"linescore,omitempty"`
			Linescore struct {
				Note                 string `json:"note"`
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
				Offense struct {
				} `json:"offense"`
				Balls   int  `json:"balls"`
				Strikes int  `json:"strikes"`
				Outs    int  `json:"outs"`
				Tie     bool `json:"tie"`
			} `json:"linescore,omitempty"`
			IsTie     bool `json:"isTie,omitempty"`
			Linescore struct {
				Note                 string `json:"note"`
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
				Offense struct {
				} `json:"offense"`
				Balls   int  `json:"balls"`
				Strikes int  `json:"strikes"`
				Outs    int  `json:"outs"`
				Tie     bool `json:"tie"`
			} `json:"linescore,omitempty"`
			Linescore struct {
				Note                 string `json:"note"`
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
				Offense struct {
					First struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"first"`
					Second struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"second"`
					Third struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"third"`
				} `json:"offense"`
				Balls   int `json:"balls"`
				Strikes int `json:"strikes"`
				Outs    int `json:"outs"`
			} `json:"linescore,omitempty"`
			Linescore struct {
				Note                 string `json:"note"`
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
				Offense struct {
				} `json:"offense"`
				Balls   int  `json:"balls"`
				Strikes int  `json:"strikes"`
				Outs    int  `json:"outs"`
				Tie     bool `json:"tie"`
			} `json:"linescore,omitempty"`
		} `json:"games"`
		Events []interface{} `json:"events"`
	} `json:"dates"`
}
