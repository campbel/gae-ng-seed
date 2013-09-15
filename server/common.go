package server

import (
	"appengine"
	"appengine/user"
	"net/http"
	"code.google.com/p/xsrftoken"
)

type XSRFEntity struct {
	Token string
}

func CreateXSRFToken(w http.ResponseWriter, r *http.Request) (string) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	// first parameter should be changed to a secret key
	token := xsrftoken.Generate("myAppsSecretToken", u.ID, "api")
	
	cookie := http.Cookie {
		Name: "token", 
		Value: token,
		Path: "/",
	}
	
	http.SetCookie(w, &cookie)

	return token
}