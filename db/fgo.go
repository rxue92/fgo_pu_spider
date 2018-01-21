package db

import (
	"fmt"
	"log"

	"github.com/rxue92/fgo_pu_spider/g"
)


func Insert2FGOPU(obj g.FGOPick) {
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
