package repos

import (
	"github.com/drum445/vehicleFinderMongo/models"
	"github.com/globalsign/mgo/bson"
)

func (db DB) GetVehicles(page int, params map[string]string) (count int, vehicles models.Vehicles) {
	// query builder, for each param add a regex lowercase search
	// to the filter
	filter := make(bson.M)
	for k, v := range params {
		if v != "" {
			filter[k] = bson.M{"$regex": v, "$options": "i"}
		}
	}

	// select from the vehicles coll using our filter and
	// decode into vehicle object
	skip := (page - 1) * 10
	limit := 10
	db.Vehicles.Find(filter).Skip(skip).Limit(limit).All(&vehicles)
	count, _ = db.Vehicles.Find(filter).Count()

	return
}

func (db DB) GetVehicle(vehicleID int) (vehicle models.Vehicle, found bool) {
	// attempt to look for vehicle, return false if it doesn't exist
	filter := bson.M{"_id": vehicleID}
	if db.Vehicles.Find(filter).One(&vehicle) == nil {
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
