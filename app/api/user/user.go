package user

import (
    "app"
    "appengine"
    "appengine/user"
    "net/http"
)

type userinfo struct {
	User *user.User
	LogoutUrl string
} 

func init() {
    http.HandleFunc("/api/user", app.CreateHandler(get, nil, nil, nil))
}

func get(w http.ResponseWriter, r *http.Request) (interface{}) {
	c := appengine.NewContext(r)
    u := user.Current(c)
    url := ""

    // true if user isn't logged in. redirect to the login url
    if u != nil {
        url, _ = user.LogoutURL(c, "/")
    }

    user := userinfo {
    	User: u,
    	LogoutUrl: url,
    }

	return user;
}