package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type Entity struct {
	Value string
}

func init() {
	http.HandleFunc("/", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	e := new(Entity)
	e.Value = "Bob"
	k := datastore.NewKey(c, "Entity", "robertsturner52@gmail.com", 0, nil)

	if _, err := datastore.Put(c, k, e); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "new=%qn", e.Value)
}