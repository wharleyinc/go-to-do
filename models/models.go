package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDo struct {
	ID        primitive.ObjectID `bson:"_id"`
	Task      string             `bson:"task"`
	Status    string             `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
