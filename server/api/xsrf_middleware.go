package api

import (
	"appengine"
	"appengine/user"
	"github.com/campbel/gore"
	"net/http"
	"server"
)

func init() {
	gore := gore.Module("api")
	gore.Middleware("/api/", xsrfHandler)
}

func xsrfHandler(w http.ResponseWriter, r *http.Request, x gore.Context) bool {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		return true
	}

	if r.Method == "GET" {
		return true
	}

	validToken := server.ValidateXSRFToken(w, r)
	if !validToken {
		x.Status(403)
		return false
	}

	return true
}
