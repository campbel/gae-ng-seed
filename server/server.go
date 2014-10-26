package server

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("server/index.html"))

type rootTemplateData struct {
	Dev bool
}

func init() {
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/logout", logoutHandle)
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {

	CreateXSRFToken(w, r)

	templates.ExecuteTemplate(w, "main", rootTemplateData{
		Dev: appengine.IsDevAppServer(),
	})
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, "/")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusFound)
}

func logoutHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u != nil {
		url, err := user.LogoutURL(c, "/")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusFound)
}
