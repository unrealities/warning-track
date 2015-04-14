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
2015/04/11/seamlb-oakmlb-1
Final
1.9
2015/04/11/detmlb-clemlb-1
Final
1.7
2015/04/11/tbamlb-miamlb-1
Final
1.7
2015/04/11/tormlb-balmlb-1
Final
0.1
2015/04/11/wasmlb-phimlb-1
Final
0
2015/04/11/nynmlb-atlmlb-1
Final
1.3
2015/04/11/pitmlb-milmlb-1
Final
0.1
2015/04/11/houmlb-texmlb-1
Final
0.4
2015/04/11/chnmlb-colmlb-1
Final
0.1
2015/04/11/lanmlb-arimlb-1
Final
0.1
2015/04/11/sfnmlb-sdnmlb-1
Final
0.1
2015/04/11/kcamlb-anamlb-1
Final
0.6
```

## TODO

* Verify leverage index (especially at the beginning of an inning)
* Rank games by current leverage index
* Cache game results
* Don't parse games that are final
* Provide links to TV/gamecast/radio
* Provide current pitcher v. hitter matchup
* Get setup on Google App engine
* Get crons setup to scan mlb live data
* Send email when leverage index crosses 3 in a game
