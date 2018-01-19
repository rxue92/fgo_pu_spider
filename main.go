package main

import (
	"flag"
	"os"

	"practice/fgospider/db"
	"practice/fgospider/g"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	db.Init()
}
