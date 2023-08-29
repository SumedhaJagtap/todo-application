package main

import (
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/todoapplication/app"
	"github.com/todoapplication/server"
)

func main() {
	log.Println("Application has started")
	os.Setenv("DBUSER", "root")
	os.Setenv("DBPASS", "josh730")

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error occured in application", r)
		}
		log.Println("Application stopped")
	}()
	app.Init()
	defer app.Close()

	server.StartServer()

}
