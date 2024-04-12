package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercise struct {
	id     primitive.ObjectID `bson:"id,omitempty"`
	name   string             `bson:"name,omitempty"`
	date   string             `bson:"date,omitempty"`
	weight float32            `bson:"weight,omitempty"`
	reps   int32              `bson:"reps,omitempty"`
	time   int32              `bson:"time,omitempty"`
}
