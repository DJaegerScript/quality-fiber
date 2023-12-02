package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "softtest.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS user (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE,
			password TEXT
		);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
