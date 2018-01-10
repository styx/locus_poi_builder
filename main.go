package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/qedus/osmpbf"
)

var nodes map[int64]*osmpbf.Node
var ways map[int64]*osmpbf.Node
var relations map[int64]*osmpbf.Relation

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Fatal("Invalid args, you must specify a PBF file")
	}

	countryName := "./" + strings.SplitN(filepath.Base(args[0]), "-", 2)[0] + ".osm.db"
	log.Println(countryName)

	_, err := copy("./template.osm.db", countryName)
	if err != nil {
		log.Fatal(err)
	}

	db := openSpatialiteDB(countryName)
	defer db.Close()

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := osmpbf.NewDecoder(f)
	err = decoder.Start(runtime.GOMAXPROCS(-1)) // use several goroutines for faster decoding
	if err != nil {
		log.Fatal(err)
	}

	process(decoder)
	store(db)
}

func store(db *sql.DB) {
	bulkInserInBatches(db, nodes, "P")
	log.Println("Nodes: DONE")

	// bulkInserInBatches(db, ways, "W")
	// log.Println("Ways: DONE")
}

func copy(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	srcFileStat, err := srcFile.Stat()
	if err != nil {
		return 0, err
	}

	if !srcFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}

	count, err := io.Copy(dstFile, srcFile)
	dstFile.Close()

	return count, err
}
