package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/drum445/vehicleFinderMongo/models"
	"github.com/drum445/vehicleFinderMongo/repos"
	"github.com/gorilla/mux"
)

type response struct {
	Count    int             `json:"count"`
	Vehicles models.Vehicles `json:"vehicles,omitempty"`
}

func GetVehicles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// default the page number to 1. If the page has been passed through as a URL
	// param, first check if it is an int, if it is use it, else reuturn 400
	page := 1
	var err error

	if req.URL.Query().Get("page") != "" {
		page, err = strconv.Atoi(req.URL.Query().Get("page"))
		if err != nil {
			http.Error(w, "page must be an int", 400)
			return
		}
	}

	// map[string]string of all our expected params
	m := map[string]string{
		"make":       req.URL.Query().Get("make"),
		"shortModel": req.URL.Query().Get("shortModel"),
		"longModel":  req.URL.Query().Get("longModel"),
		"trim":       req.URL.Query().Get("trim"),
		"derivative": req.URL.Query().Get("derivative"),
		"available":  "Y",
	}

	db := repos.Init()
	defer db.Close()
	count, vehicles := db.GetVehicles(page, m)

	// create our response object and encode to json
	var resp response
	resp.Count = count
	resp.Vehicles = vehicles
	json.NewEncoder(w).Encode(resp)
}

func GetVehicleByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	vehicleID, err := strconv.Atoi(params["vehicleID"])
	if err != nil {
		http.Error(w, "vehicle ID must be an int", 400)
		return
	}

	db := repos.Init()
	defer db.Close()
	vehicle, found := db.GetVehicle(vehicleID)

	if !found {
		http.Error(w, "vehicle ID not found", 400)
		return
	}

	vehicle.Image = repos.GetImage(vehicle.ID)
	json.NewEncoder(w).Encode(vehicle)
}

func PostVehicles(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("Vehicles.csv")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// load file then skip Header
	reader := csv.NewReader(file)
	reader.Read()

	// loop through each record create a vehicle object and append to ur vehicles
	var vehicles models.Vehicles
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

		vehicles = append(vehicles, vehicle)
	}

	// once vehicles object is created bulik insert into db
	db := repos.Init()
	defer db.Close()
	db.InsertVehicles(vehicles)

	fmt.Fprint(w, fmt.Sprintf("Finished importing %v rows", len(vehicles)))
}
