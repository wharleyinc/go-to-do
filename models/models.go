package models

import (
	"time"
)

type ToDo struct {
	// ID        primitive.ObjectID `bson:"_id"` remove so that mongodb can auto generate ids for records
	Task      string    `bson:"task"`
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
