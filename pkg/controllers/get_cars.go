package controllers

import (
	"context"
	"errors"

	"bitbucket.org/8BitsKW/go-backend/pkg/database"
	"bitbucket.org/8BitsKW/go-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCars() ([]models.Car, error) {

	cursor, err := database.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, errors.New("error while getting cars")
	}
	var results []models.Car
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}
