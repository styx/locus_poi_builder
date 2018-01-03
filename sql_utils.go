package main

import (
	"database/sql"
	"log"

	sqlite3 "github.com/mattn/go-sqlite3"
)

func runQuery(db *sql.DB, query string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func openSpatialiteDB(dbPath string) *sql.DB {
	sql.Register("sqlite3_with_spatialite",
		&sqlite3.SQLiteDriver{
			Extensions: []string{"mod_spatialite"},
		})

	db, err := sql.Open("sqlite3_with_spatialite", dbPath)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	return db
}
