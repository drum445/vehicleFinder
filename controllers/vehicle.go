package controllers

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/drum445/vehicleFinder/models"
	"github.com/drum445/vehicleFinder/repos"
	"github.com/gorilla/mux"
)

type response struct {
	Count    int             `json:"count"`
	Vehicles models.Vehicles `json:"vehicles,omitempty"`
}

func GetVehicles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// map[string]string of all our expected params
	m := make(map[string]string)
	m["make"] = req.URL.Query().Get("make")
	m["shortModel"] = req.URL.Query().Get("shortModel")
	m["longModel"] = req.URL.Query().Get("longModel")
	m["trim"] = req.URL.Query().Get("trim")
	m["derivative"] = req.URL.Query().Get("derivative")
	m["available"] = "Y"

	db := repos.Init()
	vehicles := db.GetVehicles(m)
	defer db.Close()

	var resp response
	resp.Count = len(vehicles)
	resp.Vehicles = vehicles
	json.NewEncoder(w).Encode(resp)
}

func GetVehicleByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	vehicleID, err := strconv.Atoi(params["vehicleID"])
	if err != nil {
		http.Error(w, "vehicle ID must be an int", 401)
		return
	}

	var vehicle models.Vehicle
	vehicle.ID = vehicleID

	db := repos.Init()
	found := db.GetVehicle(&vehicle)
	defer db.Close()

	if !found {
		http.Error(w, "vehicle ID not found", 401)
		return
	}
	json.NewEncoder(w).Encode(vehicle)
}

func PostVehicles(w http.ResponseWriter, req *http.Request) {
	// open db connection and close after func end
	db := repos.Init()
	defer db.Close()

	file, err := os.Open("Vehicles.csv")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// skip Header
	reader.Read()

	// loop through each record create a vehicle object
	// and import it into mongo
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		var vehicle models.Vehicle
		vehicle.ID, _ = strconv.Atoi(record[0])
		vehicle.Make = record[1]
		vehicle.ShortModel = record[2]
		vehicle.LongModel = record[3]
		vehicle.Trim = record[4]
		vehicle.Derivative = record[5]
		vehicle.Introduced = record[6]
		vehicle.Discontinued = record[7]
		vehicle.Available = record[8]

		// Insert into DB
		db.InsertVehicle(vehicle)
	}
}
