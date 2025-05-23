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
	createusersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err := Db.Exec(createusersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	_, err = Db.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	_, err = Db.Exec(createRegistrationsTable)
	if err != nil {
		panic(err)
	}
}
