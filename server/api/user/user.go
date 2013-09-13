package user

import (
	"appengine"
	"appengine/user"
	"server/gore"
)

func init() {
	gore.Get("/api/user/", get)
}

func get(req *gore.Request, res *gore.Response) {
	c := appengine.NewContext(req.Raw)
	u := user.Current(c)
	res.Send(u)
}
