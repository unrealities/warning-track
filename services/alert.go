package services

import (
	"strconv"

	"github.com/unrealities/warning-track/models"
)

func AlertMessage(a models.Alert) string {
	teams := Teams()
	homeTeam := "NA"
	awayTeam := "NA"
	homeTeamHashtag := ""
	awayTeamHashtag := ""
	for _, t := range teams {
		if a.Teams.Home == t.Id {
			homeTeam = t.Abbr
			homeTeamHashtag = t.Hashtag
		}
		if a.Teams.Away == t.Id {
			awayTeam = t.Abbr
			awayTeamHashtag = t.Hashtag
		}
	}

	return awayTeam + " " + strconv.Itoa(a.Score.Away) + " - " +
		homeTeam + " " + strconv.Itoa(a.Score.Home) + ". " +
		a.HalfInning + " " + strconv.Itoa(a.Inning) + ". " +
		OutsMessage(a.Outs) + ". " +
		BaseRunnerStateMessage(a.BaseRunnerState) + ". " +
		BattingMessage(a.HalfInning, a.Batter, awayTeamHashtag, homeTeamHashtag) +
		". mlb-tv: " + a.Link
}

// 0:none; 1:1b; 2:2b; 3:3b; 4:1b,2b; 5:1b,3b; 6:2b,3b; 7:1b,2b,3b
func BaseRunnerStateMessage(b int) string {
	switch b {
	case 0:
		return "Bases empty"
	case 1:
		return "Runner on 1st"
	case 2:
		return "Runner on 2nd"
	case 3:
		return "Runner on 3rd"
	case 4:
		return "Runners on 1st,2nd"
	case 5:
		return "Runners on 1st,3rd"
	case 6:
		return "Runners on 2nd,3rd"
	case 7:
		return "Bases loaded"
	}
	return "Base status unknown"
}

func OutsMessage(o int) string {
	switch o {
	case 0:
		return "No outs"
	case 1:
		return "1 out"
	case 2:
		return "2 outs"
	case 3:
		return "End of inning"
	}
	return "Unknown outs"
}

func BattingMessage(halfInning, batter, awayHashtag, homeHashtag string) string {
	if halfInning == "Top" {
		return "#" + awayHashtag + " " + batter + " at-bat vs #" + homeHashtag
	} else {
		return "#" + homeHashtag + " " + batter + " at-bat vs #" + awayHashtag
	}

}
