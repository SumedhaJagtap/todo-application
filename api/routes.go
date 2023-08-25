package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", defaultHandler).Methods(http.MethodGet)
	return
}
