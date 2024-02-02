package db

import (
	"context"

	"github.com/ashiqsabith123/chat-svc/pkg/config"
	logs "github.com/ashiqsabith123/love-bytes-proto/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase(config config.Config) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.MongoDB.MongoUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("chat-svc").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		logs.ErrLog.Println(err)
		return nil, err
	}
	logs.GenLog.Println("Mongo connected successfully", result)

	return client, nil

}
