package api

import (
	"appengine"
	"appengine/user"
	"github.com/campbel/gore"
	"server"
)

func init() {
	gore := gore.Module("api")
	gore.Middleware("/api/", xsrfHandler)
}

func xsrfHandler(req *gore.Request, res *gore.Response) bool {
	c := appengine.NewContext(req.Raw)
	u := user.Current(c)

	if u == nil {
		return true
	}

	if req.Raw.Method == "GET" {
		return true
	}

	validToken := server.ValidateXSRFToken(res.Raw, req.Raw)
	if !validToken {
		res.Status(403)
		return false
	}

	return true
}
