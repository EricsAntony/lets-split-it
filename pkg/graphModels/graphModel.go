package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func InitDB() {
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	db, err := sql.Open("mysql", "root:root@tcp(localhost)/expensetracker")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
 		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}
