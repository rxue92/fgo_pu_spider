package g

import "time"

type FGOPick struct {
	RId       string `json:"rid"`
	Sname     string `json:"sname"`
	Rname     string `json:"rname"`
	Servant   string `json:"servant"`
	Rarity    int    `json:"rarity"`
	Timestamp int64  `json:"timestamp"`
	MainUp5   bool   `json:"main_up_5"`
	MainUp4   bool   `json:"main_up_4"`
	MinorUp5  bool   `json:"minor_up_5"`
}

type PickUpMod struct {
	MainSsr  string   `json:"mainSsr"`
	MinorSsr string   `json:"minorSsr"`
	MainSr   []string `json:"mainSr"`
}

// 默认四星非up从者
func NewFGOPick() *FGOPick {
	return &FGOPick{
		Rarity:    4,
		Timestamp: time.Now().Unix(),
		MainUp5:   false,
		MainUp4:   false,
		MinorUp5:  false,
	}
}
