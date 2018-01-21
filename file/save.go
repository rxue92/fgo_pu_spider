package file

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/rxue92/fgo_pu_spider/g"
)

func Save2CSV(data []g.FGOPick) {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal("create file fail")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Println(err)
		}
	}
}
