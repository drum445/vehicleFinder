package repos

import (
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	Conn *sql.DB
}

func (db *DB) Init() {
	var err error
	db.Conn, err = sql.Open("mysql", "root:password@/")
	// db.Conn, err = sql.Open("mysql", "user:password@tcp(server.com)/")

	db.Conn.Exec("USE vehicle_finder")

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func (db DB) CreateDB() {
	db.Conn.Exec("CREATE DATABASE IF NOT EXISTS vehicle_finder")
	db.Conn.Exec(`CREATE TABLE IF NOT EXISTS vehicle_finder.vehicle (
		vehicle_id int(11) NOT NULL,
		make varchar(45) DEFAULT NULL,
		short_model varchar(45) DEFAULT NULL,
		long_model varchar(45) DEFAULT NULL,
		trim varchar(45) DEFAULT NULL,
		derivative varchar(45) DEFAULT NULL,
		introduced varchar(45) DEFAULT NULL,
		discontinued varchar(45) DEFAULT NULL,
		available varchar(45) DEFAULT NULL,
		PRIMARY KEY (vehicle_id)
	  ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
	  `)
}

func (db DB) Close() {
	db.Conn.Close()
}
