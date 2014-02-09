package server

import (
	"appengine"
	"appengine/user"
	"code.google.com/p/xsrftoken"
	"net/http"
)

type XSRFEntity struct {
	Token string
}

func CreateXSRFToken(w http.ResponseWriter, r *http.Request) string {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		return ""
	}

	// first parameter should be changed to a secret key
	token := xsrftoken.Generate("myAppsSecretToken", u.ID, "api")

	cookie := http.Cookie{
		Name:  "XSRF-TOKEN",
		Value: token,
		Path:  "/",
	}

	http.SetCookie(w, &cookie)

	return token
}
