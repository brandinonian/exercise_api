package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init_DB_Connection() (*mongo.Client, error) {
	// Set your MongoDB connection string
	connection_string := "mongodb://localhost:27017"

	// Create a ClientOptions object with your connection string
	client_options := options.Client().ApplyURI(connection_string)

	// Create a client
	client, err := mongo.Connect(context.TODO(), client_options)
	if err != nil {
		return nil, err
	}

	// Ping the server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Close the database connection
func Disconnect_DB(client *mongo.Client) error {
	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}

	return nil
}
