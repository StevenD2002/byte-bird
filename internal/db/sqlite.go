package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitSQLiteDB initializes the SQLite database connection
func InitSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to SQLite database")
	return db, nil
}
