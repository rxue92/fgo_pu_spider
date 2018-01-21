package rw

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/json-iterator/go"
)

type BlData struct {
	Id    string `json:"id"`
	Sname string `json:"sname"`
	Rname string `json:"rname"`
	Info  string `json:"info"`
	Star  string `json:"star"`
}

type BLResp struct {
	Code int    `json:"code"`
	Data []BlData `json:"data"`
}

var (
	gameId  = 112
	gameKey = "a5f36e53ab3b0c41"
	rows    = 31
)

func GetData() ([]BlData, error) {
	// main post: http://game.bilibili.com/fgo/news.html#!news/1/6/1509

	url := "http://activity.biligame.com/board/list?game_id=112&game_key=a5f36e53ab3b0c41&rows=31"
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return []BlData{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return []BlData{}, err
	}

	var resv BLResp
	err = jsoniter.Unmarshal(body, &resv)
	if err != nil {
		log.Println(err)
		return []BlData{}, err
	}

	//log.Println(resv.Data)

	return resv.Data, nil

}
