package handler

import (
	"api/internal/database"
	"api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GET /exercises/:date
func Get_Exercises_By_Date(ctx *gin.Context) {

	date := ctx.Param("date")

	cursor, err := database.Exercises.Find(ctx, bson.M{"date": date})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var exercises []model.Exercise
	if err := cursor.All(ctx, &exercises); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exercises)
}

// POST /exercises
func Add_Exercise(ctx *gin.Context) {
	var body model.CreateExerciseRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.Exercises.InsertOne(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add exercise"})
		return
	}

	exercise := model.Exercise{
		Id:     res.InsertedID.(primitive.ObjectID),
		Name:   body.Name,
		Date:   body.Date,
		Weight: body.Weight,
		Reps:   body.Reps,
		Time:   body.Time,
	}

	ctx.JSON(http.StatusCreated, exercise)
}

// DELETE /exercises/:id
func Delete_Exercise(ctx *gin.Context) {
	id := ctx.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := database.Exercises.DeleteOne(ctx, bson.M{"_id": _id})

	if res.DeletedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "exercise not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "exercise deleted"})
}
