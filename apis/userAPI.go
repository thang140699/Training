package apis

import (
	"context"
	"encoding/json"
	"mongo-with-golang/entities"
	"mongo-with-golang/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func FindUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["Id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "missing")
		return
	}
	user, err := models.FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

//connect mongo and getall dataa
func GetAll(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// mongoclient, err = mongo.Connect(ctx, client)
	if err != nil {
		panic(err)
	}
	defer cancel()
	//check
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	collec := client.Database("domain").Collection("web")
	cursor, err := collec.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var result []bson.M
	if err := cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	responseWithJSON(response, http.StatusOK, result)
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.CreateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "could not create")
			return
		}
		responseWithJSON(response, http.StatusOK, "successfully")
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["Id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "missing URL")
		return
	}
	result := models.DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusOK, "Could not delete")
	}
	responseWithJSON(response, http.StatusOK, "Delete")
}
func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}
func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
