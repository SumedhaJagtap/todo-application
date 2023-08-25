package api

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println("Default handler started")
	defer log.Println("Default handler ended")

	select {
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	default:
		fmt.Fprint(w, "Redirect to default page")
	}
}
