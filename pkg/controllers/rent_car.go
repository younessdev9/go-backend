package controllers

import (
	"context"
	"errors"
	"time"

	"bitbucket.org/8BitsKW/go-backend/pkg/database"
	"bitbucket.org/8BitsKW/go-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func RentCar(registration string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	filter := bson.M{"registration": registration}
	var car models.Car
	err := database.Collection.FindOne(ctx, filter).Decode(&car)

	if err != nil {
		return "", errors.New("car does not exist")
	}

	if car.Rented {
		return "", errors.New("car is already rented")
	} else {
		update := bson.M{
			"$set": bson.M{"rented": true},
		}
		var car models.Car
		err := database.Collection.FindOneAndUpdate(ctx, filter, update).Decode(&car)
		if err != nil {
			return "", err
		}

		return "car successfully rented", nil
	}
}
