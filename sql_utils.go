package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	sqlite3 "github.com/mattn/go-sqlite3"
	"github.com/qedus/osmpbf"
	"gopkg.in/cheggaaa/pb.v2"
)

const batchSize = 200

var valueStrings [batchSize]string

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

func bulkInsertInBatches(db *sql.DB, rows map[int64]*osmpbf.Node, nodeType string, progressBar *pb.ProgressBar) {
	bufNode := make([]*osmpbf.Node, 0, batchSize)

	idx := 0

	foldersCounter := 0
	foldersValueStrings := make([]string, 0, batchSize)
	foldersValueArgs := make([]interface{}, 0, batchSize*3)

	tagsCounter := 0
	tagsValueStrings := make([]string, 0, batchSize)
	tagsValueArgs := make([]interface{}, 0, batchSize*3)

	for i := 0; i < batchSize; i++ {
		valueStrings[i] = "(?, ?, ?, GeomFromText(?, 4326))"
	}

	for _, node := range rows {
		progressBar.Add(1)

		validForInsert, folders := tagsToFolders(node.Tags)

		if validForInsert {
			idx++
			bufNode = append(bufNode, node)

			for _, folder := range *folders {
				foldersCounter++
				foldersValueStrings = append(foldersValueStrings, "(?, ?, ?)")
				foldersValueArgs = append(foldersValueArgs, idx)
				foldersValueArgs = append(foldersValueArgs, (*folder)[0])
				foldersValueArgs = append(foldersValueArgs, (*folder)[1])
			}

			for tag, tagVal := range node.Tags {
				if tagKeysWhiteList.Contains(tag) {
					tagsCounter++
					tagsValueStrings = append(tagsValueStrings, "(?, ?, ?)")
					tagsValueArgs = append(tagsValueArgs, idx)
					tagsValueArgs = append(tagsValueArgs, getTagKeyID(db, tag))
					tagsValueArgs = append(tagsValueArgs, getTagValueID(db, tagVal))
				}
			}

			// Store buffers: Node
			if idx%batchSize == 0 {
				bulkInsertPoint(db, bufNode, nodeType)

				bufNode = make([]*osmpbf.Node, 0, batchSize)
			}

			// Store buffers: Folders
			if foldersCounter >= batchSize {
				tx, err := db.Begin()
				if err != nil {
					log.Fatal(err)
				}

				stmt := fmt.Sprintf(
					"INSERT INTO Points_Root_Sub (Points_id, FoldersRoot_id, FoldersSub_id) VALUES %s",
					strings.Join(foldersValueStrings, ","))

				_, err = tx.Exec(stmt, foldersValueArgs...)
				if err != nil {
					log.Fatal(err)
				}

				tx.Commit()

				foldersCounter = 0
				foldersValueStrings = make([]string, 0, batchSize)
				foldersValueArgs = make([]interface{}, 0, batchSize*3)
			}

			// Store buffers: Tags
			if tagsCounter >= batchSize {
				tx, err := db.Begin()
				if err != nil {
					log.Fatal(err)
				}

				stmt := fmt.Sprintf(
					"INSERT INTO Points_Key_Value (Points_id, TagKeys_id, TagValues_id) VALUES %s",
					strings.Join(tagsValueStrings, ","))

				_, err = tx.Exec(stmt, tagsValueArgs...)
				if err != nil {
					log.Fatal(err)
				}

				tx.Commit()

				tagsCounter = 0
				tagsValueStrings = make([]string, 0, batchSize)
				tagsValueArgs = make([]interface{}, 0, batchSize*3)
			}
		}
	}
}

func bulkInsertPoint(db *sql.DB, unsavedRows []*osmpbf.Node, nodeType string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	valueArgs := make([]interface{}, 0, batchSize*4)

	for _, row := range unsavedRows {
		valueArgs = append(valueArgs, row.ID)
		valueArgs = append(valueArgs, nodeType)
		valueArgs = append(valueArgs, row.Tags["name"])
		valueArgs = append(valueArgs, fmt.Sprintf("POINT(%.6f %.6f)", row.Lon, row.Lat))
	}

	stmt := fmt.Sprintf(
		"INSERT INTO Points (id, type, name, geom) VALUES %s",
		strings.Join(valueStrings[0:len(unsavedRows)-1], ","),
	)

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

func getTagKeyID(db *sql.DB, tagName string) int64 {
	rows, err := db.Query("SELECT id FROM TagKeys WHERE name = ? LIMIT 1;", tagName)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id sql.NullInt64
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	if id.Valid {
		return id.Int64
	}

	result, err := db.Exec("INSERT INTO TagKeys (name) VALUES (?)", tagName)
	if err != nil {
		log.Fatal(err)
	}

	lastID, _ := result.LastInsertId()
	return lastID
}

func getTagValueID(db *sql.DB, tagValue string) int64 {
	rows, err := db.Query("SELECT id FROM TagValues WHERE name = ? LIMIT 1", tagValue)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id sql.NullInt64
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	if id.Valid {
		return id.Int64
	}

	result, err := db.Exec("INSERT INTO TagValues (name) VALUES (?)", tagValue)
	if err != nil {
		log.Fatal(err)
	}

	lastID, _ := result.LastInsertId()
	return lastID
}
