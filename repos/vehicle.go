package repos

import (
	"github.com/drum445/vehicleFinder/models"
	"gopkg.in/mgo.v2/bson"
)

func (db DB) GetVehicles(params map[string]string) (vehicles models.Vehicles) {
	// query builder, for each param add a regex lowercase search
	// to the filter
	filter := bson.M{}
	for k, v := range params {
		if v != "" {
			filter[k] = bson.M{"$regex": v, "$options": "i"}
		}
	}

	// select from the vehicles coll using our filter and
	// decode into vehicle object
	db.Vehicles.Find(filter).All(&vehicles)

	return
}

func (db DB) GetVehicle(vehicleID int) (vehicle models.Vehicle, found bool) {
	// attempt to look for vehicle, return false if it doesn't exist
	filter := bson.M{"_id": vehicleID}
	err := db.Vehicles.Find(filter).One(&vehicle)
	if err == nil {
		found = true
	}

	return
}

func (db DB) InsertVehicle(vehicle models.Vehicle) {
	db.Vehicles.Insert(vehicle)
}

func (db DB) InsertVehicles(vehicles models.Vehicles) {
	s := make([]interface{}, len(vehicles))
	for i, v := range vehicles {
		s[i] = v
	}

	x := db.Vehicles.Bulk()
	x.Insert(s...)
	x.Run()
}
