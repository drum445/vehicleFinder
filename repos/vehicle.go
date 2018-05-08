package repos

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/drum445/vehicleFinder/models"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

type VehicleRepo struct {
	DB
}

func NewVehicleRepo() (vr *VehicleRepo) {
	vr.Init()
	return
}

func rowToVehicle(row *sql.Row) (v models.Vehicle, err error) {
	err = row.Scan(&v.ID, &v.Make, &v.ShortModel, &v.LongModel, &v.Trim, &v.Derivative, &v.Introduced, &v.Discontinued, &v.Available)
	return
}

func rowsToVehicle(rows *sql.Rows) (vehicles models.Vehicles) {
	for rows.Next() {
		var v models.Vehicle
		rows.Scan(&v.ID, &v.Make, &v.ShortModel, &v.LongModel, &v.Trim, &v.Derivative, &v.Introduced, &v.Discontinued, &v.Available)
		vehicles = append(vehicles, v)
	}

	return
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
	vr.Conn.QueryRow("SELECT COUNT(vehicle_id) FROM vehicle WHERE "+strings.Join(columns, " AND "), values...).Scan(&count)

	// add the limit and skip values to our array then get the results
	limit := 10
	skip := (page - 1) * limit

	values = append(values, skip, limit)
	rows, _ := vr.Conn.Query("SELECT * FROM vehicle WHERE "+strings.Join(columns, " AND ")+" LIMIT ?, ?", values...)
	vehicles = rowsToVehicle(rows)

	return
}

func (vr VehicleRepo) GetVehicle(vehicleID int) (vehicle models.Vehicle, found bool) {
	// attempt to look for vehicle, return false if it doesn't exist
	row := vr.Conn.QueryRow("SELECT * FROM vehicle WHERE vehicle_id = ?", vehicleID)
	vehicle, err := rowToVehicle(row)
	if err == nil {
		found = true
	}

	return
}

func (vr VehicleRepo) InsertVehicle(v models.Vehicle) {
	vr.Conn.Exec("INSERT INTO vehicle VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		v.ID, v.Make, v.ShortModel, v.LongModel, v.Trim, v.Derivative, v.Introduced, v.Discontinued, v.Available)
}
