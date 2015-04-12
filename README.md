# warning-track
parse mlb data and retrieve alerts

Currently returns:
* List of current games
* Game status
* Current leverage index

## How To
```go
$GOPATH/bin/warning-track

2015/04/11/bosmlb-nyamlb-1
Final
0.2
2015/04/11/slnmlb-cinmlb-1
Final
0.2
2015/04/11/minmlb-chamlb-1
Final
1.5
```

## TODO

* Rank games by current leverage index
* Cache game results
* Don't parse games that are final
* Provide links to TV/gamecast/radio
* Provide current pitcher v. hitter matchup
* Get setup on Google App engine
* Get crons setup to scan mlb live data
* Send email when leverage index crosses 3 in a game
