package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDb() {
	var err error
	Db, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
	createTable()
}

func createTable() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)`
	_, err := Db.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}
}
