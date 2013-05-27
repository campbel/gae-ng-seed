package app

import (
    "appengine"
    "appengine/user"
    "net/http"
    "html/template"
)


var templates = template.Must(template.ParseGlob("app/index.html"))

func init() {
    http.HandleFunc("/login", loginHandle)
	// http.Handle("/", http.FileServer(http.Dir("./app")))
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)

    if u == nil {
        url, err := user.LoginURL(c, "/")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        templates.ExecuteTemplate(w, "index", map[string]string {
            "login": url,
        })
    } else {
        url, err := user.LogoutURL(c, "/")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        templates.ExecuteTemplate(w, "index", map[string]string {
            "logout": url,
            "email": u.Email,
        })
    }
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    c := appengine.NewContext(r)
    u := user.Current(c)

    // true if user isn't logged in. redirect to the login url
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
    
    // otherwise redirect to the root
    w.Header().Set("Location", "/")
    w.WriteHeader(http.StatusFound)
}