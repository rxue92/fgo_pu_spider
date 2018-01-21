package main

import (
	"flag"
	"os"

	//"github.com/rxue92/fgo_pu_spider/db"
	"github.com/rxue92/fgo_pu_spider/g"
	"github.com/rxue92/fgo_pu_spider/rw"
)

func main() {
	cfg := flag.String("c", "cfg.example.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	//db.Init()

	// TODO: 改成在配置文件中配置up从者等信息
	rw.InitParams(false)

	rw.Loop(360)
}
