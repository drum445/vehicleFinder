package repos

import (
	"fmt"
	"strings"

	"github.com/drum445/vehicleFinder/models"
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
		if v != "" {
			values = append(values, fmt.Sprintf("%%%v%%", v))
			columns = append(columns, fmt.Sprintf("%s like ?", k))
		}
	}

	// get the count of our query
	vr.conn.Select(&count, "SELECT COUNT(vehicle_id) FROM vehicle WHERE "+strings.Join(columns, " AND "), values...)

	// add the limit and skip values to our array then get the results
	limit := 10
	skip := (page - 1) * limit

	values = append(values, skip, limit)
	vr.conn.Select(&vehicles, "SELECT * FROM vehicle WHERE "+strings.Join(columns, " AND ")+" LIMIT ?, ?", values...)

	return
}

func (vr VehicleRepo) GetVehicle(vehicleID int) (vehicle models.Vehicle, found bool) {
	// attempt to look for vehicle, return false if it doesn't exist
	err := vr.conn.Select(&vehicle, "SELECT * FROM vehicle WHERE vehicle_id = ?", vehicleID)
	if err == nil {
		found = true
	}

	return
}

func (vr VehicleRepo) InsertVehicle(v models.Vehicle) {
	vr.conn.NamedExec(`INSERT INTO vehicle VALUES
						(:id, :make, :short_model, :long_model, :trim, :derivative, :introduced, :discontinued, :available)`, v)
}
