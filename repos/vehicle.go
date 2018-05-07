package repos

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/drum445/vehicleFinder/models"
	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

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

func (db DB) GetVehicles(page int, params map[string]string) (count int, vehicles models.Vehicles) {
	// query builder, for each kvp in our params map make a key like %value%
	var values []interface{}
	var columns []string
	for k, v := range params {
		if v != "" {
			values = append(values, fmt.Sprintf("%%%v%%", v))
			columns = append(columns, fmt.Sprintf("%s like ?", k))
		}
	}

	rows, _ := db.Conn.Query("SELECT * FROM vehicle WHERE "+strings.Join(columns, " AND "), values...)
	vehicles = rowsToVehicle(rows)

	db.Conn.QueryRow("SELECT COUNT(vehicle_id) FROM vehicle WHERE "+strings.Join(columns, " AND "), values...).Scan(&count)
	return
}

func (db DB) GetVehicle(vehicleID int) (vehicle models.Vehicle, found bool) {
	// attempt to look for vehicle, return false if it doesn't exist
	row := db.Conn.QueryRow("SELECT * FROM vehicle WHERE vehicle_id = ?", vehicleID)
	vehicle, err := rowToVehicle(row)
	if err == nil {
		found = true
	}

	return
}

func (db DB) InsertVehicle(v models.Vehicle) {
	db.Conn.Exec("INSERT INTO vehicle VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)",
		v.ID, v.Make, v.ShortModel, v.LongModel, v.Trim, v.Derivative, v.Introduced, v.Discontinued, v.Available)
}
