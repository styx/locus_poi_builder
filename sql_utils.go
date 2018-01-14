package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	sqlite3 "github.com/mattn/go-sqlite3"
	"github.com/qedus/osmpbf"
)

const batchSize = 200

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

func bulkInserInBatches(db *sql.DB, rows map[int64]*osmpbf.Node, nodeType string) {
	buf := make([]*osmpbf.Node, 0, batchSize)

	idx := 0
	for _, node := range rows {
		idx++
		buf = append(buf, node)
		if idx%batchSize == 0 {
			bulkInsertPoint(db, buf, nodeType)

			idx = 0
			buf = make([]*osmpbf.Node, 0, batchSize)
		}
	}
}

func bulkInsertPoint(db *sql.DB, unsavedRows []*osmpbf.Node, nodeType string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	valueStrings := make([]string, 0, batchSize)
	valueArgs := make([]interface{}, 0, batchSize*4)

	// if hasTags(node.Tags) {
	// 	folders := tagsToFolders(node.Tags)
	// 	for _, arr := range *folders {
	// 		fmt.Printf("%d\t%d %d\n", node.ID, arr[0], arr[1])
	// 	}
	// }

	for _, row := range unsavedRows {
		valueStrings = append(valueStrings, "(?, ?, ?, GeomFromText(?, 4326))")
		valueArgs = append(valueArgs, row.ID)
		valueArgs = append(valueArgs, nodeType)
		valueArgs = append(valueArgs, row.Tags["name"])
		valueArgs = append(valueArgs, fmt.Sprintf("POINT(%.6f %.6f)", row.Lon, row.Lat))
	}

	stmt := fmt.Sprintf("INSERT INTO Points (id, type, name, geom) VALUES %s", strings.Join(valueStrings, ","))

	_, err = tx.Exec(stmt, valueArgs...)
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

	return db
}

func query(db *sql.DB) {
	q := "SELECT id, type, name, AsText(geom) FROM Points ORDER BY name DESC LIMIT 20;"
	rows, err := db.Query(q)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var v1, v2, v3, v4 string
		if err = rows.Scan(&v1, &v2, &v3, &v4); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s %s %s\n", v1, v2, v3, v4)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
