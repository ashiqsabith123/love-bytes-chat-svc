package db

import (
	"context"

	"github.com/ashiqsabith123/chat-svc/pkg/config"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	Chat_rooms *mongo.Collection
}

func ConnectToDatabase(config config.Config) *Collections {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MongoDB.MongoUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := client.Database("chat-svc").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		logs.ErrLog.Fatal(err)
		return nil
	}
	logs.GenLog.Println("Mongo connected successfully", result)

	chatDatabase := client.Database("chat-svc")

	chatDatabase.CreateCollection(context.Background(), "chat_rooms")

	chat_rooms := chatDatabase.Collection("chat_rooms")

	return &Collections{
		Chat_rooms: chat_rooms,
	}

}
