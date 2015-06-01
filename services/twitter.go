package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mrjones/oauth"
	"github.com/unrealities/warning-track/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
)

type Twitter struct {
	consumer    *oauth.Consumer
	accessToken *oauth.AccessToken
}

func GetTwitterCredentials(w http.ResponseWriter, r *http.Request) models.Credentials {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Credentials") //.Filter("Key Name =", "Twitter")
	tc := models.Credentials{}
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

func NewTwitter(tc models.Credentials, r *http.Request) *Twitter {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	twitter := new(Twitter)
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

func (t *Twitter) Post(url string, params map[string]string) (interface{}, error) {
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
