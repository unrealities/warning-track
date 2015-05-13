# warning-track
Parse mlb data and retrieve alerts based on [leverage index](http://www.fangraphs.com/library/misc/li/)

[http://warningtrack.co/](http://warningtrack.co/)
[@warningtrackco](http://twitter.com/warningtrackco)

## TODO

* Verify leverage index (especially at the beginning of an inning)
* Smarter "fetchAllStatuses": Within so many minutes of a scheduled game, we should try to update that game's status. Or manually set their status to In Progress to allow it to be caught in the refresh cycle.
