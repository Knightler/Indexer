package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() error {
	client, err := mongo.NewClient(options.Client().ApplyURL("mongodb+srv://<user>:<pass>@cluster0.xxxx.mongodb.net"))
	if err != nil { return err }
	ctx, cancel := context.WithTimeout(context.Background()m 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil { return err }
	DB = client.Database("ai_database")
	return nil
}
