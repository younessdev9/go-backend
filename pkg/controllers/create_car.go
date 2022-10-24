package controllers

import (
	"context"
	"fmt"
	"time"

	"bitbucket.org/8BitsKW/go-backend/pkg/database"
	"bitbucket.org/8BitsKW/go-backend/pkg/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCar(car *models.Car) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	validator := validator.New()
	if err := validator.Struct(car); err != nil {
		return "", fmt.Errorf("please make sure that all required field are present")
	}
	filter := bson.M{"registration": car.Registration}
	count, err := database.Collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return "", fmt.Errorf("registration already exist, registration : %s ", car.Registration)
	} else {
		result, err := database.Collection.InsertOne(ctx, car)
		if err != nil {
			return "", fmt.Errorf("error while creating with registration : %s ", car.Registration)
		}
		stringObjectID := result.InsertedID.(primitive.ObjectID).Hex()

		return stringObjectID, nil
	}
}
