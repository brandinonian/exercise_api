package main

import (
	"api/internal/database"
	"api/internal/handler"
	"fmt"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/exercises/:date", handler.Get_Exercises_By_Date)
	r.POST("/exercises", handler.Add_Exercise)
	r.DELETE("/exercises/:id", handler.Delete_Exercise)

	r.Run(":8080")
}
