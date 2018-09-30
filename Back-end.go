package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func init() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/twitter/register", registerTwitterUser)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

type userName struct {
	Value   string
	Claimed bool
}

func registerTwitterUser(w http.ResponseWriter, r *http.Request) {
	//First we need to read the twitter username/ID
	//send the username to the database
	//Send the code to claim prize

	ctx := appengine.NewContext(r)
	twitterID := r.URL.Query().Get("ID")
	k := datastore.NewKey(ctx, "userName", twitterID, 0, nil)
	e := new(userName)

	// if err := datastore.Get(ctx, k, e); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	if _, err := datastore.Put(ctx, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set(twitterID, "text/plain; charset=utf-8")

}
