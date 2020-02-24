# Stats API MLB
# Description
For a long time I used mlb's legacy api. But they have shut that down and moved to
statsapi.mlb.com
I will try to document the endpoints I need for warning-track here.

## Host
statsapi.mlb.com

## Endpoints
### Schedule
#### URI
/api/v1/schedule?sportId=1&date=MM/DD/YYYY

#### Example Response
```
{
  "copyright" : "Copyright 2020 MLB Advanced Media, L.P.  Use of any content on this page acknowledges agreement to the terms posted here http://gdx.mlb.com/components/copyright.txt",
  "totalItems" : 16,
  "totalEvents" : 0,
  "totalGames" : 16,
  "totalGamesInProgress" : 14,
  "dates" : [ {
    "date" : "2020-02-24",
    "totalItems" : 16,
    "totalEvents" : 0,
    "totalGames" : 16,
    "totalGamesInProgress" : 14,
    "games" : [ {
      "gamePk" : 606569,
      "link" : "/api/v1.1/game/606569/feed/live",
      "gameType" : "S",
      "season" : "2020",
      "gameDate" : "2020-02-24T18:05:00Z",
      "status" : {
        "abstractGameState" : "Live",
        "codedGameState" : "I",
        "detailedState" : "In Progress",
        "statusCode" : "I",
        "abstractGameCode" : "L"
      },
      "teams" : {
        "away" : {
          "leagueRecord" : {
            "wins" : 0,
            "losses" : 2,
            "pct" : ".000"
          },
          "score" : 2,
          "team" : {
            "id" : 121,
            "name" : "New York Mets",
            "link" : "/api/v1/teams/121"
          },
          "splitSquad" : false,
          "seriesNumber" : 4
        },
        "home" : {
          "leagueRecord" : {
            "wins" : 1,
            "losses" : 1,
            "pct" : ".500"
          },
          "score" : 1,
          "team" : {
            "id" : 120,
            "name" : "Washington Nationals",
            "link" : "/api/v1/teams/120"
          },
          "splitSquad" : false,
          "seriesNumber" : 3
        }
      },
      "venue" : {
        "id" : 5000,
        "name" : "FITTEAM Ballpark of the Palm Beaches",
        "link" : "/api/v1/venues/5000"
      },
      "content" : {
        "link" : "/api/v1/game/606569/content"
      },
      "gameNumber" : 1,
      "publicFacing" : true,
      "doubleHeader" : "N",
      "gamedayType" : "Y",
      "tiebreaker" : "N",
      "calendarEventID" : "14-606569-2020-02-24",
      "seasonDisplay" : "2020",
      "dayNight" : "day",
      "scheduledInnings" : 9,
      "inningBreakLength" : 0,
      "gamesInSeries" : 1,
      "seriesGameNumber" : 1,
      "seriesDescription" : "Spring Training",
      "recordSource" : "S",
      "ifNecessary" : "N",
      "ifNecessaryDescription" : "Normal Game"
    }, ... ],
    "events" : [ ]
  } ]
}
```

### BoxScore
#### URI
api/v1/game/${gamepk_id}/boxscore

#### Comments
I don't think this provides me with real time events. Too much data.

### TODO
Need a "master scoreboard" and a "game state"
