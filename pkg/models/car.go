package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Model        string             `json:"model" bson:"model" validate:"required"`
	Mileage      string             `json:"mileage" bson:"milage" validate:"required"`
	Rented       bool               `json:"rented" bson:"rented"`
	Registration string             `json:"registration" bson:"registration" validate:"required"`
}
