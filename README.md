# warning-track
parse mlb data and retrieve alerts

## How To
`warning-track {url_to_game_events_json_file}`

i.e.

```go
$GOPATH/bin/warning-track http://gdb.com/components/game/mlb/year_2015/month_04/day_10/gid_2015_04_10_detmlb_clemlb_1/game_events.json
0.6
```

## TODO

* Determine all current active games
* Rank games by current leverage index
* Provide links to TV/gamecast/radio
* Provide current pitcher v. hitter matchup
* Get setup on Google App engine
* Get crons setup to scan mlb live data
* Send email when leverage index crosses 3 in a game
