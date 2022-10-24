package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bitbucket.org/8BitsKW/go-backend/pkg/controllers"
	"bitbucket.org/8BitsKW/go-backend/pkg/models"
)

func CarHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("should return cars list"))
}
func AddCarHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var car models.Car
	err := decoder.Decode(&car)
	if err != nil {
		panic(err)
	}
	err = controllers.CreateCar(&car)

	response := models.ResponseObject{Message: err.Error()}

	jsonResponse, jsonError := json.Marshal(response)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else {
		jsonResponse, jsonError = json.Marshal(models.ResponseObject{Message: "car created"})

		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		w.Write([]byte("car created successfully"))
	}
}
