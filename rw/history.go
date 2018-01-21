package rw

import (
	"log"
	"strconv"
	"time"

	"github.com/rxue92/fgo_pu_spider/g"
)

var (
	LastRid      = 0
	PickUps      *g.PickUpMod
	checkPickUps bool
	ssrCount = 0
	srCount = 0
)

func InitParams(check bool,  ) {
	checkPickUps = check
}



func trimData(picks []BlData) []BlData {
	var ret []BlData

	for _, pick := range picks {
		if num, _ := strconv.Atoi(pick.Id); num > LastRid {
			LastRid = num
			ret = append(ret, pick)
			continue
		}
		//log.Printf("lastRid:%d, current rid:%s", LastRid, pick.Id)
	}

	return ret
}

// 默认4星非up从者
func convertData(pick BlData) *g.FGOPick {
	ret := g.NewFGOPick()
	ret.Servant = pick.Info
	ret.RId = pick.Id
	ret.Sname = pick.Sname
	ret.Rname = pick.Rname
	if pick.Star == "SSR" {
		ssrCount++
		ret.Rarity = 5
	}
	srCount++

	pickUPCheck(ret)

	return ret
}

func pickUPCheck(pick *g.FGOPick) {
	if !checkPickUps {
		return
	}

	if pick.Rarity == 4 {
		for _, name := range PickUps.MainSr {
			if pick.Servant == name {
				pick.MainUp4 = true
				return
			}
		}

		pick.MainUp4 = false
		return
	}

	pick.MainUp4 = false
	if pick.Servant == PickUps.MainSsr {
		pick.MainUp5 = true
		return
	}

	if pick.Servant == PickUps.MinorSsr {
		pick.MinorUp5 = true
		return
	}

	return

}

func Loop(limit int) {
	count := 0
	for count < limit {
		data, err := GetData()
		if err != nil {
			log.Println(err)
			continue
		}
		trimmed := trimData(data)
		for _, data := range trimmed {
			pick := convertData(data)
			log.Printf("%+v\n", pick)
		}

		count++
		time.Sleep(30 * time.Second)
	}

	log.Printf("SSR:%d, SR:%d, Ratio:%f", ssrCount, srCount, float32(srCount)/float32(ssrCount))
}
