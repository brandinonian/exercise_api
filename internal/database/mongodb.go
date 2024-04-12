package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(db string) error {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	localClient, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}
	client = localClient

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	return err

}

func Close() error {
	return client.Disconnect(context.Background())
}
