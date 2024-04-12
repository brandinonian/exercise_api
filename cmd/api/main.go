package main

import (
	"api/internal/database"
	"fmt"
	"net/http"
)

func main() {
	if err := database.Init("exercises"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database connected...")

	defer func() {
		if err := database.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /exercises", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("POST /exercises", func(w http.ResponseWriter, r *http.Request) {

	})
}
