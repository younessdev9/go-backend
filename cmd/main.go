package main

import (
	"net/http"
	"os"

	"bitbucket.org/8BitsKW/go-backend/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/", routes.CarHandler).Methods("GET")

	http.ListenAndServe(":"+port, router)
}
