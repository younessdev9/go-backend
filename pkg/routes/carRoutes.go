package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"bitbucket.org/8BitsKW/go-backend/pkg/controllers"
	"bitbucket.org/8BitsKW/go-backend/pkg/models"
	"github.com/gorilla/mux"
)

func GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	results, err := controllers.GetAllCars()
	if err != nil {
		jsonResponse, jsonError := json.Marshal(models.ResponseObject{Message: err.Error()})
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(jsonResponse)
	} else {
		jsonResponse, jsonError := json.Marshal(results)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
func AddCarHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var car models.Car
	err := decoder.Decode(&car)
	if err != nil {
		panic(err)
	}
	id, err := controllers.CreateCar(&car)
	if err != nil {
		response := models.ResponseObject{Message: err.Error()}
		jsonResponse, jsonError := json.Marshal(response)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else {
		jsonResponse, jsonError := json.Marshal(models.CreatedCar{Id: id})
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
func RentCarHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message, err := controllers.RentCar(vars["registration"])
	if err != nil {
		response := models.ResponseObject{Message: err.Error()}
		jsonResponse, jsonError := json.Marshal(response)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else {
		jsonResponse, jsonError := json.Marshal(models.ResponseObject{Message: message})
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
func ReturnCarHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var returnReq models.ReturnCarRequest
	err := decoder.Decode(&returnReq)
	switch {
	case err == io.EOF:
		response := models.ResponseObject{Message: "please provide request body"}
		jsonResponse, jsonError := json.Marshal(response)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
		return
	case err != nil:
		panic(err)
	}

	message, err := controllers.ReturnCar(vars["registration"], returnReq)
	if err != nil {
		response := models.ResponseObject{Message: err.Error()}
		jsonResponse, jsonError := json.Marshal(response)
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResponse)
	} else {
		jsonResponse, jsonError := json.Marshal(models.ResponseObject{Message: message})
		if jsonError != nil {
			fmt.Println("Unable to encode JSON")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
