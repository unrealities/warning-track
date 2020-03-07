package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/unrealities/warning-track/models"
	"github.com/unrealities/warning-track/services"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
)

func GameJSON(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)

	liveGames := []models.Game{}
	liveStatuses := []models.Status{}

	_, get_cache_err := memcache.JSON.Get(c, "Game", &liveGames)
	if get_cache_err != nil && get_cache_err != memcache.ErrCacheMiss {
		http.Error(w, get_cache_err.Error(), http.StatusInternalServerError)
	}
	if get_cache_err != nil {
		q := datastore.NewQuery("Game")

		_, Err := q.GetAll(c, &liveGames)
		if Err != nil {
			http.Error(w, Err.Error(), http.StatusInternalServerError)
			return
		}

		item := &memcache.Item{
			Key:    "Game",
			Object: liveGames,
		}
		setErr := memcache.JSON.Set(c, item)
		if setErr != nil {
			http.Error(w, setErr.Error(), http.StatusInternalServerError)
		}
	}

	warningTrackGames := make([]models.WtGame, len(liveGames))

	_, get_cache_err = memcache.JSON.Get(c, "Status", &liveStatuses)
	if get_cache_err != nil && get_cache_err != memcache.ErrCacheMiss {
		http.Error(w, get_cache_err.Error(), http.StatusInternalServerError)
	}
	if get_cache_err != nil {
		q := datastore.NewQuery("Status")

		_, Err := q.GetAll(c, &liveStatuses)
		if Err != nil {
			http.Error(w, Err.Error(), http.StatusInternalServerError)
			return
		}

		item := &memcache.Item{
			Key:    "Status",
			Object: liveStatuses,
		}
		setErr := memcache.JSON.Set(c, item)
		if setErr != nil {
			http.Error(w, setErr.Error(), http.StatusInternalServerError)
		}
	}

	for k, lg := range liveGames {
		warningTrackGames[k].Id = lg.Id
		warningTrackGames[k].Teams = lg.Teams
		warningTrackGames[k].Links = lg.Links

		t, _ := time.Parse("2006/01/02 3:04PM -0700", lg.DateTime)
		warningTrackGames[k].DateTime = t.UTC().Format(time.RFC3339)

		for _, ls := range liveStatuses {
			if ls.GameId == lg.Id {
				warningTrackGames[k].Status = ls
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := JSONMarshal(warningTrackGames, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func SetGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)
	// If before 12pm UTC (8am EST). Display the results from the day before
	gameTime := time.Now().UTC()
	if gameTime.Hour() < 12 {
		gameTime = time.Now().UTC().Add(-12 * time.Hour)
	}

	games := []models.Game{}
	msb := services.MasterScoreboard(gameTime, r)
	for _, s := range msb.Dates[0].Games {
		log.Printf("msb.Dates[0].Games: %+v", s)
		g := models.Game{}
		g.Id = s.GamePk
		g.Teams.Away = s.Teams.Away.Team.ID
		g.Teams.Home = s.Teams.Home.Team.ID
		g.DateTime = s.GameDate.Format("2006-01-02") // Might need to parse: 2020-02-24T18:05:00Z

		for _, e := range s.Content.Media.Epg {
			log.Printf("s.Content.Media.Epg: %+v", e)
			if e.Title == "MLBTV" && len(e.Items) > 0 {
				log.Printf("MLBTV e.Items: %+v", e.Items)
				// TODO: this may be a dangerous assumption that the first item has the contentID we want
				g.Links.MlbTv = services.MlbApiMlbTvLinkToUrl(s.GamePk, e.Items[0].ContentID)
			}
			break
		}

		games = append(games, g)
	}

	// store games
	item := &memcache.Item{
		Key:    "Game",
		Object: games,
	}
	setErr := memcache.JSON.Set(c, item)
	if setErr != nil {
		http.Error(w, setErr.Error(), http.StatusInternalServerError)
	}

	keys := make([]*datastore.Key, len(games))
	for k := range keys {
		keys[k] = datastore.NewKey(c, "Game", "", int64(games[k].Id), nil)
	}

	_, err := datastore.PutMulti(c, keys, games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
