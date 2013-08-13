package user

import (
	"appengine"
	"appengine/user"
	"github.com/campbel/gore"
)

func init() {
	gore.CreateHandler("/api/user/").Get(get)
}

func get(r *gore.Request) interface{} {
	c := appengine.NewContext(r.Request)
	u := user.Current(c)
	return u
}
