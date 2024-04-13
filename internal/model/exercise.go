package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	Id     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	Date   string             `bson:"date,omitempty" json:"date,omitempty"`
	Weight float32            `bson:"weight,omitempty" json:"weight,omitempty"`
	Reps   int32              `bson:"reps,omitempty" json:"reps,omitempty"`
	Time   int32              `bson:"time,omitempty" json:"time,omitempty"`
}

type CreateExerciseRequest struct {
	Name   string  `bson:"name,omitempty" json:"name,omitempty"`
	Date   string  `bson:"date,omitempty" json:"date,omitempty"`
	Weight float32 `bson:"weight,omitempty" json:"weight,omitempty"`
	Reps   int32   `bson:"reps,omitempty" json:"reps,omitempty"`
	Time   int32   `bson:"time,omitempty" json:"time,omitempty"`
}
