package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoPassword := os.Getenv("MONGODB_PASSWORD")
	clientOptions := options.Client().ApplyURI("mongodb+srv://car-rental:" + mongoPassword + "@cluster0.2rchchr.mongodb.net/?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database("cars-rental").Collection("cars")

}
func Setup() {
	fmt.Println("Connection to db")
}
