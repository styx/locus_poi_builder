package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/deckarep/golang-set"
	"github.com/qedus/osmpbf"
	"gopkg.in/cheggaaa/pb.v2"
)

var nodes map[int64]*osmpbf.Node
var ways map[int64]*osmpbf.Node
var relations map[int64]*osmpbf.Relation

var tagKeysWhiteList mapset.Set

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

	loadWhiteList()

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

	processPbf(decoder)
	store(db)
}

func store(db *sql.DB) {
	barTemplate := pb.ProgressBarTemplate(`{{red "Points:"}} {{bar . "[" "\u2588" "\u25B6" " " "]" | green}} {{percent . | yellow}}`)
	bar := barTemplate.Start(len(nodes))

	bulkInsertInBatches(db, nodes, "P", bar)
	bar.Finish()
	log.Println("Nodes: DONE")

	barTemplate = pb.ProgressBarTemplate(`{{red "Ways:"}} {{bar . "[" "\u2588" "\u25B6" " " "]" | green}} {{percent . | yellow}}`)
	bar = barTemplate.Start(len(ways))

	bulkInsertInBatches(db, ways, "W", bar)
	bar.Finish()
	log.Println("Ways: DONE")
}
