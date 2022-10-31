package uploadfile

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"strconv"
	"time"
)

func Upload() {
	pathfile := strconv.Itoa(time.Now().Year()) + "/" + time.Now().Month().String() + "/" + strconv.Itoa(time.Now().Day()) + "/Domain.txt"
	file, err := os.Open(pathfile)
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//for scanner.Scan() {
	//
	//	fmt.Println(scanner.Text())
	//}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
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
		newBson := bson.D{{Key: "Year", Value: time.Now().Year()}, {Key: "Month", Value: time.Now().Month()}, {Key: "Day", Value: time.Now().Day()}, {Key: "Domain", Value: line}}
		docs = append(docs, newBson)
		uploadDomians, err := uploadDomain.InsertOne(ctx, newBson)
		if err != nil {
			log.Fatalln("False ", uploadDomians)
		}
	}

}
