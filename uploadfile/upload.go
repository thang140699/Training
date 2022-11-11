package uploadfile

import (
	"bufio"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"strconv"
	"time"
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	//mongoclient, err = mongo.Connect(ctx, client)
	if err != nil {
		panic(err)
	}
	defer cancel()
	//check
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	uploadDomain := client.Database("domain").Collection("web")

	var docs []interface{}
	for scanner.Scan() {
		line := scanner.Text()
		newBson := setTime{time.Now(), line}
		docs = append(docs, newBson)
		uploadDomians, err := uploadDomain.InsertOne(ctx, newBson)
		if err != nil {
			log.Fatalln("False ", uploadDomians)
		}
	}
	return ""

}
