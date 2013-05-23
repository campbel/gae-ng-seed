package user

import (
    "app"
    "appengine"
    "appengine/user"
    "net/http"
)

func init() {
    http.HandleFunc("/api/user", app.CreateHandler(get, nil, nil, nil))
}

func get(w http.ResponseWriter, r *http.Request) (interface{}) {
	c := appengine.NewContext(r)
    u := user.Current(c)
	return u;
}