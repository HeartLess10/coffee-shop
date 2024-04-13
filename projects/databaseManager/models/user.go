package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User[key string | primitive.ObjectID] struct {
	Id     key    `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Age    int    `json:"age" bson:"age"`
	Email  string `json:"email" bson:"email"`
	UserID string `json:"userId" bson:"userId"`
}
