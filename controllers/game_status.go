package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unrealities/warning-track/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
)

func DeleteGamesStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)

	//delete games
	g := []models.Game{}
	gq := datastore.NewQuery("Game").KeysOnly()

	gk, Err := gq.GetAll(c, &g)
	if Err != nil {
		http.Error(w, Err.Error(), http.StatusInternalServerError)
		return
	}

	datastore.DeleteMulti(c, gk)
	memcache.Delete(c, "Game")

	//delete statuses
	s := []models.Status{}
	sq := datastore.NewQuery("Status").KeysOnly()

	sk, Err := sq.GetAll(c, &s)
	if Err != nil {
		http.Error(w, Err.Error(), http.StatusInternalServerError)
		return
	}

	datastore.DeleteMulti(c, sk)
	memcache.Delete(c, "Status")

	//reload current games and statuses
	SetGames(w, r, nil)
	SetAllStatuses(w, r, nil)
}
