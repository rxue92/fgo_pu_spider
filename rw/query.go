package rw

type BlData struct {
	Id    int    `json:"id"`
	Sname string `json:"sname"`
	Rname string `json:"rname"`
	Info  string `json:"info"`
	Star  string `json:"star"`
}

var (
	gameId  = 112
	gameKey = "a5f36e53ab3b0c41"
	rows    = 31
)

func getData() {

}
