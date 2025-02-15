package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() *sql.DB {
	var err error
	db, err := sql.Open("sqlite3", "data/test.db")
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS note (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT NOT NULL,
      content TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatal("Failed to create tables:", err)
	}
	return db
}
