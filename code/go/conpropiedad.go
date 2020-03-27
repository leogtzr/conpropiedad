package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

var (
	dbWords    []Word
	wordFormat = *regexp.MustCompile(`^.+;.+;.+$`)
)

func main() {
	tag := flag.String("tag", "", "tag: the tag to search for ... ")
	db := flag.String("db", "", "db file")
	flag.Parse()

	if !validateOptions(*tag, *db) {
		log.Fatal("error: missing tag options")
	}

	words, err := loadInputFile(*db, &wordFormat)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(words))

}
