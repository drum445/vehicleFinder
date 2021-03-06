package repos

import (
	"fmt"
	"strings"

	"github.com/drum445/vehicleFinder/backend/models"
)

type VehicleRepo struct {
	DB
}

func NewVehicleRepo() *VehicleRepo {
	var vr VehicleRepo
	vr.Init()
	return &vr
}

func (vr VehicleRepo) GetVehicles(page int, params map[string]string) (count int, vehicles models.Vehicles) {
	// query builder, for each kvp in our params map make a key like %value%
	var values []interface{}
	var columns []string
	for k, v := range params {
		if k == "free" {
			values = append(values, fmt.Sprintf("%%%v%%", v))
			columns = append(columns, "CONCAT(vehicle_id, make, short_model, long_model, trim, derivative) like ?")
		} else if v != "" {
			values = append(values, fmt.Sprintf("%%%v%%", v))
			columns = append(columns, fmt.Sprintf("%s like ?", k))
		}
	}

	// get the count of our query
	vr.conn.Get(&count, "SELECT COUNT(vehicle_id) FROM vehicle WHERE "+strings.Join(columns, " AND "), values...)

	// add the limit and skip values to our array then get the results
	limit := 10
	skip := (page - 1) * limit

	values = append(values, skip, limit)
	vr.conn.Select(&vehicles, "SELECT * FROM vehicle WHERE "+strings.Join(columns, " AND ")+" LIMIT ?, ?", values...)

	return
}

func (vr VehicleRepo) GetVehicle(vehicleID int) (vehicle models.Vehicle, found bool) {
	// attempt to look for vehicle, return false if it doesn't exist
	err := vr.conn.Get(&vehicle, "SELECT * FROM vehicle WHERE vehicle_id = ?;", vehicleID)
	if err == nil {
		found = true
	}

	return
}

func (vr VehicleRepo) InsertVehicle(v models.Vehicle) {
	vr.conn.Exec("INSERT IGNORE INTO vehicle VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", v.ID, v.Make, v.ShortModel, v.LongModel, v.Trim, v.Derivative,
		v.Introduced, v.Discontinued, v.Available)
}
