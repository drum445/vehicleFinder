package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Disable strictslash so both "/api/vehicle" and "/api/vehicle/" work
	router.StrictSlash(true)
	router.HandleFunc("/api/vehicle/", getVehicles).Methods("GET")
	router.HandleFunc("/api/vehicle/{vehicleID}/", getVehicleByID).Methods("GET")
	router.HandleFunc("/api/vehicle/import", postVehicles).Methods("POST")

	fmt.Println("started on port :5000")

	// set our handlers and start the server
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(router)))
}
