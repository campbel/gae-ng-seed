package common

import (
	"encoding/json"
	"net/http"
)

// Create a Handler that maps to the given functions
func CreateHandler(get, post, put, del func(http.ResponseWriter, *http.Request) (v interface{})) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var obj interface{}
		switch r.Method {
		case "GET":
			obj = get(w, r)
		case "POST":
			obj = post(w, r)
		case "PUT":
			obj = put(w, r)
		case "DELETE":
			obj = del(w, r)
		}
		buff, err := json.Marshal(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(buff)
		}
		return
	}
}
