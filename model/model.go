package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movies struct {
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
}
