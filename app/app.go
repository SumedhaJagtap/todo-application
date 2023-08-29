package app

import "github.com/todoapplication/db"

func Init() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}
}

func Close() {
	db.Close()
}
