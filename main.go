package main

import (
	"database/sql"
	"flag"
	"fmt"
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
	showBlacklistedTags := flag.Bool("show-blacklisted-tags", false, "Shows tags found in PBF but not listed in "+tagKeysWhiteListFileName)
	flag.BoolVar(showBlacklistedTags, "sbt", false, "Shows tags found in PBF but not listed in "+tagKeysWhiteListFileName)

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		flag.PrintDefaults()
		log.Fatal("Invalid args, you must specify a PBF file")
	}

	countryName := "./" + strings.SplitN(filepath.Base(args[0]), "-", 2)[0] + ".osm.db"
	log.Println(countryName)

	loadWhiteList()

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

	if *showBlacklistedTags {
		printBlacklistedTags()
		os.Exit(0)
	}

	_, err = copy("./template.osm.db", countryName)
	if err != nil {
		log.Fatal(err)
	}

	db := openSpatialiteDB(countryName)
	defer db.Close()
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

func printBlacklistedTags() {
	tagKeysBlackList := mapset.NewSet()

	for _, node := range nodes {
		validForInsert, _ := tagsToFolders(node.Tags)

		if validForInsert {
			for tagKey := range node.Tags {
				if !tagKeysWhiteList.Contains(tagKey) {
					tagKeysBlackList.Add(tagKey)
				}
			}
		}
	}

	for _, node := range ways {
		for tagKey := range node.Tags {
			if !tagKeysWhiteList.Contains(tagKey) {
				tagKeysBlackList.Add(tagKey)
			}
		}
	}

	it := tagKeysBlackList.Iterator()
	for tagName := range it.C {
		fmt.Println(tagName.(string))
	}
}
