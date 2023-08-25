package api

import (
	"log"
	"net/http"
)

func StartServer() {
	router := initRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}
	log.Fatal(srv.ListenAndServe())
}
