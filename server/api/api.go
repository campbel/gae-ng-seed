package api

import (
	"github.com/campbel/gore"
	"net/http"
)

func init() {
	gore := gore.Module("api")
	http.HandleFunc("/api/", gore.Handler)
}
