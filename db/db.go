package db

import (
	"database/sql"
	"log"

	"github.com/rxue92/fgo_pu_spider/g"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	initPortaldb()
}

func initPortaldb() {
	var err error
	DB, err = sql.Open("postgres", g.Config().Postgresql.Addr)
	if err != nil {
		log.Fatalln("open postgresql fail:", err)
	}

	DB.SetMaxIdleConns(g.Config().Postgresql.MaxIdle)

	err = DB.Ping()
	if err != nil {
		log.Fatalln("ping postgresql fail:", err)
	}
}
