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

func (db DB) GetVehicle(vehicle *models.Vehicle) bool {
	// attempt to look for vehicle, return false if it doesn't exist
	filter := bson.M{"_id": vehicle.ID}
	err := db.Vehicles.Find(filter).One(&vehicle)
	if err != nil {
		return false
	}

	// call the image API and assign our image to the vehicle
	vehicle.Image = models.GetImage(vehicle.ID)
	return true
}

func (db DB) InsertVehicle(vehicle models.Vehicle) {
	db.Vehicles.Insert(vehicle)
}
