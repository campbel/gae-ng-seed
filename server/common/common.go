package common

import (
	"encoding/json"
	"net/http"
)

// Shortcut for routing http methods to handlers and serializing structs to JSON
func CreateHandler(get, post, put, del func(http.ResponseWriter, *http.Request) (v interface{})) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var obj interface{}

		/* 
			MUX the incomming requests method
			responde with Status '501 - StatusNotImplemented' for any methods we don't have handlers for
		*/
		switch r.Method {
		case "GET":
			if get != nil {
				obj = get(w, r)
			} else {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}
		case "POST":
			if post != nil {
				obj = post(w, r)
			} else {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}
		case "PUT":
			if put != nil {
				obj = put(w, r)
			} else {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}
		case "DELETE":
			if del != nil {
				obj = del(w, r)
			} else {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}
		default:
			w.WriteHeader(http.StatusNotImplemented)
			return
		}

		/*
			Handlers might not give us data as they are also capable of sending responses. If we have nothing return.
			If no response is written, Status 200 - OK will be returned.
		*/
		if obj != nil {
			buff, err := json.Marshal(obj)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write(buff)
			}
		}

		return
	}
}