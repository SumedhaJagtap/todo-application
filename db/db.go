package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() (err error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "todo",
		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	log.Println("DB connected")
	return
}

func Close() {
	db.Close()
}

func GetDB() *sql.DB {
	return db
}
