package utility

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func Success(status int, response interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func Error(status int, response interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
