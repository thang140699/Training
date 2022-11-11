package localhost

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	//"mongo-with-golang/controller"
	controller "mongo-with-golang/controller"
	"mongo-with-golang/services"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroll   controller.Controll
	ctx            context.Context
	err            error
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connect established")
	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.(usercollection, ctx)
	usercontroll = controller.New(userservice)
	server = gin.Default()
}
func Connect() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	controller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":8000"))
}
