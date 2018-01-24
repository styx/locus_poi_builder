package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/deckarep/golang-set"
)

const tagKeysWhiteListFileName = "tag_keys_white_list.txt"

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
		return 0, fmt.Errorf("%s is not a regular file", tagKeysWhiteListFileName)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}

	count, err := io.Copy(dstFile, srcFile)
	dstFile.Close()

	return count, err
}

func loadWhiteList() {
	tagKeysWhiteList = mapset.NewSet()

	srcFile, err := os.Open(tagKeysWhiteListFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	srcFileStat, err := srcFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if !srcFileStat.Mode().IsRegular() {
		log.Fatalf("%s is not a regular file", tagKeysWhiteListFileName)
	}

	scanner := bufio.NewScanner(srcFile)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			tagKeysWhiteList.Add(line)
		}
	}
}
