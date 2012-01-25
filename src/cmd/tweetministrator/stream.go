package main

import (
	"net/http"
	"encoding/json"
)

func LatestTweets(user string) (t []Tweet, e error) {
	r, e := http.DefaultClient.Get("https://api.twitter.com/1/statuses/user_timeline.json?screen_name="+user)
	if e != nil {
		return nil, e
	}

	dec := json.NewDecoder(r.Body)
	e = dec.Decode(&t)
	return
}

