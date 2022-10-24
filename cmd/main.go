package main

import (
	"net/http"
	"os"

	"bitbucket.org/8BitsKW/go-backend/pkg/database"
	"bitbucket.org/8BitsKW/go-backend/pkg/routes"
	"github.com/gorilla/mux"
)

func init() {
	database.Setup()
}
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/cars", routes.GetCarsHandler).Methods("GET")
	router.HandleFunc("/cars", routes.AddCarHandler).Methods("POST")
	router.HandleFunc("/cars/{registration}/rentals", routes.RentCarHandler).Methods("POST")

	http.ListenAndServe(":"+port, router)
}
