package main

import (
	"flag"
	"log"
)

func main() {
	tagShort := flag.String("t", "", "tag: the tag to search for ... ")
	tagLong := flag.String("tag", "", "tag: the tag to search for ... ")
	flag.Parse()

	if !validateOptions(*tagShort) && !validateOptions(*tagLong) {
		log.Fatal("error: missing tag options")
	}
}
