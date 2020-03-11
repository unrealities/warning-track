## General Architecture

### Backend

- Written in Go (GoLang) to work with Google App Engine (GAE)
- Design choices on how data is stored were made specifically to stay free on GAE
- Jobs run on a cron (app/cron.yaml) to fetch general game information and status
- Data is fetched from MLB's undocumented/unsupported API
- When interesting games occur, information is pushed to twitter (@warningtrackco)
- Static data is stored in JSON files (see: 'Static Data')
- `app/controllers`: Logic for taking requests and returning data from services
- `app/models`: Logic for structs and how data is organized
- `app/services`: Logic for business data
- `app/routers`: Logic for mapping of routes

### Static Data

- `app/base_out.json`: stores all the possible combinations of runners on base and the number of outs.
- `app/game_state.json`: stores all the possible game states that we care about. (If a team is winning/losing by more than 4 runs in any inning).
- `app/team.json`: stores general mlb team data. Including what to include as their hashtag in tweets.

### Fronted

- Written in Angular.js (1.x.x) to be a lightweight consumer of the GO API
- Code is contained in `app/www`
- No CSS frameworks are used.
- MLB Team logos are the property of MLB and accurate as of 2015.
- Other images are created by myself and use svg when possible
- `www/lib/svg4everybody.min.js` is used to improve svg resolution when scaling

## Getting Setup

- [Download and install Go on your machine](https://golang.org/dl/)
- `go get github.com/unrealities/warning-track/app`
- `go install github.com/unrealities/warning-track/app`
- [Download and install Google Cloud SDK](https://dl.google.com/dl/cloudsdk/channels/rapid/GoogleCloudSDKInstaller.exe)
- `gcloud init`

## Deploying New Code

- If `go install` doesn't work, can do `go get -u github.com/unrealities/warning-track/app`
- `gcloud config set project warning-track`
- `gcloud app deploy --version 4-1-0` from `/warning-track/app` directory

## GO Version

- go112+ does not support `login:admin` in `app.yaml` 
