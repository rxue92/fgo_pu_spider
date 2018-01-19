package db

import (
	"fmt"
	"log"
)

type FGOPickUps struct {
	RId       int    `json:"rid"`
	Sname     string `json:"sname"`
	Rname     string `json:"rname"`
	Servant   string `json:"servant"`
	Rarity    int    `json:"rarity"`
	Timestamp int    `json:"timestamp"`
}

func Insert2FGOPU(obj FGOPickUps) {
	sql := fmt.Sprintf(`
	insert into fgo_up (
		rid,
		sname,
		rname,
		servant,
		rarity,
		creat_at,
		main_up5,
		main_up4
	) values (
		%d, '%s', '%s', '%s', '%d', %d, %s, %s
	`,
		obj.RId,
		obj.Sname,
		obj.Rname,
		obj.Servant,
		obj.Rarity,
		obj.Timestamp,
		"true",
		"false",
	)

	res, err := DB.Exec(sql)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.RowsAffected())
	return
}
