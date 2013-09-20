package server

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
	"github.com/campbel/gore"
)

var templates = template.Must(template.ParseGlob("server/index.html"))

type TemplateData struct {
  Dev bool
  XSRFToken string
}

func init() {
	http.HandleFunc("/login", loginHandle)
	http.HandleFunc("/logout", logoutHandle)
	gore.Start("/api/")
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "main", TemplateData{
      Dev: appengine.IsDevAppServer(),
      XSRFToken: CreateXSRFToken(w, r),
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
