package controllers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"bitbucket.org/8BitsKW/go-backend/pkg/database"
	"bitbucket.org/8BitsKW/go-backend/pkg/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func ReturnCar(registration string, request models.ReturnCarRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	filter := bson.M{"registration": registration}
	var car models.Car
	err := database.Collection.FindOne(ctx, filter).Decode(&car)

	if err != nil {
		return "", errors.New("car does not exist")
	}
	validator := validator.New()

	if err := validator.Struct(request); err != nil {
		return "", errors.New("please provide kilometersDriven field")
	}

	if err != nil {
		fmt.Errorf("error while converting milage to int %s", err.Error())
	}

	if !car.Rented {
		return "", errors.New("car is not yet rented")
	} else {
		currentCarMilage, err := strconv.Atoi(car.Mileage)
		if err != nil {
			return "", errors.New("error while parsing KilometersDriven")
		}
		newMilage := strconv.Itoa(currentCarMilage + request.KilometersDriven)
		update := bson.M{
			"$set": bson.M{"milage": newMilage, "rented": false},
		}
		var car models.Car
		err = database.Collection.FindOneAndUpdate(ctx, filter, update).Decode(&car)
		if err != nil {
			return "", err
		}

		return "car is now available", nil
	}
}
