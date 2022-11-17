package uploadfile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type setTime struct {
	Time   time.Time `json:"Time" bson:"Time"`
	Domain string    `json:"Domain" bson:"Domain"`
}

var (
	_          = godotenv.Load(".env")
	domainFile = os.Getenv("domainTxT")
	url        = os.Getenv("url")
	domainZip  = os.Getenv("domainZip")
	setime     []setTime
)

func Upload(Time time.Time) string {
	pathfile := strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day()) + domainFile

	file, err := os.Open(pathfile)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// 	connect
	// https: //domain:9000/api

	// settime
	// var docs []interface{}
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	newBson := setTime{time.Now(), line}
	// 	docs = append(docs, newBson)
	// 	uploadDomians, err := uploadDomain.InsertOne(ctx, newBson)
	// 	if err != nil {
	// 		log.Fatalln("False ", uploadDomians)
	// 	}
	// }
	return ""

}
