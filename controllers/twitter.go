package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unrealities/warning-track/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// Only to be used manually to update Twitter Credentials.
// Never store any keys, tokens or secrets in the code.
func SetTwitterCredentials(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)

	t := &models.Credentials{
		ConsumerKey:       "",
		ConsumerSecret:    "",
		AccessToken:       "",
		AccessTokenSecret: ""}

	k := datastore.NewKey(c, "Credentials", "Twitter", 0, nil)

	_, err := datastore.Put(c, k, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
