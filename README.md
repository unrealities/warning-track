# warning-track
parse mlb data and retrieve alerts

Currently returns list of in progress games order by leverage index descending

## How To
```go
$GOPATH/bin/warning-track

2015/04/13/miamlb-atlmlb-1: 5.5
2015/04/13/nyamlb-balmlb-1: 3.7
2015/04/13/tbamlb-tormlb-1: 1.9
2015/04/13/milmlb-slnmlb-1: 1.9
2015/04/13/detmlb-pitmlb-1: 1.5
2015/04/13/colmlb-sfnmlb-1: 0.6
2015/04/13/phimlb-nynmlb-1: 0.5
2015/04/13/anamlb-texmlb-1: 0.2
2015/04/13/kcamlb-minmlb-1: 0.1
2015/04/13/wasmlb-bosmlb-1: 0.1
2015/04/13/oakmlb-houmlb-1: 0.1
2015/04/13/arimlb-sdnmlb-1: 0.1
2015/04/13/cinmlb-chnmlb-1: 0.0
2015/04/13/seamlb-lanmlb-1: 0.0
```

## TODO

* Verify leverage index (especially at the beginning of an inning)
* Cache game results
* Don't parse games that are final
* Provide links to TV/gamecast/radio
* Provide current pitcher v. hitter matchup
* Get setup on Google App engine
* Get crons setup to scan mlb live data
* Send email when leverage index crosses 3 in a game
