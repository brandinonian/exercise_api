package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	fmt.Println("Connecting to database...")

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connected to database...")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Println(err)
		}
	}()

	// db_client, err := data.Init_DB_Connection()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// mux := http.NewServeMux()

	// mux.HandleFunc("GET /exercises/{date}", func(w http.ResponseWriter, r *http.Request) {
	// 	date := r.PathValue("date")
	// 	fmt.Fprintf(w, "%s", date)
	// })

	// mux.HandleFunc("POST /exercises", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("POST hit...")
	// })

	// if err := http.ListenAndServe("localhost:8080", mux); err != nil {
	// 	fmt.Println("Listening on 8080...")
	// 	fmt.Println(err.Error())
	// }

	// defer data.Disconnect_DB(db_client)
}
