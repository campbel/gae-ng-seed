package user

import (
	"appengine"
	"appengine/user"
	"net/http"
	"server/common"
)

func init() {
	http.HandleFunc("/api/user", common.CreateHandler(get, nil, nil, nil))
}

func get(w http.ResponseWriter, r *http.Request) interface{} {
	c := appengine.NewContext(r)
	u := user.Current(c)
	return u
}
