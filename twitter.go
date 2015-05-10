package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrjones/oauth"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
)

type twitter struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
}

type credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetTwitterCredentials(w http.ResponseWriter, r *http.Request) credentials {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Credentials") //.Filter("Key Name =", "Twitter")
	tc := credentials{}
	t := q.Run(c)
	for {
		_, err := t.Next(&tc)
		if err == datastore.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	return tc
}

// Only to be used manually to update Twitter Credentials.
// Never store any keys, tokens or secrets in the code.
func SetTwitterCredentials(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := appengine.NewContext(r)

	t := &credentials{
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

func NewTwitter(tc credentials, r *http.Request) *twitter {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	twitter := new(twitter)
	twitter.consumer = oauth.NewCustomHttpClientConsumer(
		tc.ConsumerKey,
		tc.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
		client)
	twitter.accessToken = &oauth.AccessToken{tc.AccessToken, tc.AccessTokenSecret, nil}
	return twitter
}

func (t *twitter) Post(url string, params map[string]string) (interface{}, error) {
	response, err := t.consumer.Post(url, params, t.accessToken)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// decode
	var result interface{}
	err = json.Unmarshal(b, &result)
	return result, err
}

func Tweet(s string, w http.ResponseWriter, r *http.Request) {
	tc := GetTwitterCredentials(w, r)
	twitter := NewTwitter(tc, r)

	_, err := twitter.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		map[string]string{"status": s})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
