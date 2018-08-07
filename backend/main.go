package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/drum445/vehicleFinder/backend/controllers"
	"github.com/drum445/vehicleFinder/backend/repos"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// open connection, make sure db/table exist else create them then close conn
	var db repos.DB
	db.Init()
	db.CreateDB()
	db.Close()

	router := mux.NewRouter()

	vehicleRouter := router.PathPrefix("/api/vehicle").Subrouter()
	vehicleRouter.HandleFunc("", controllers.GetVehicles).Methods("GET")
	vehicleRouter.HandleFunc("", controllers.PostVehicles).Methods("POST")
	vehicleRouter.HandleFunc("/{vehicleID}", controllers.GetVehicleByID).Methods("GET")

	fmt.Println("started on port :5000")

	// set our handlers and start the server
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "PUT", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(suffixMiddleware(router))))
}
