package repos

import (
	"github.com/globalsign/mgo"
)

type DB struct {
	Session  *mgo.Session
	Vehicles *mgo.Collection
}

func Init() *DB {
	var db DB
	db.Session, _ = mgo.Dial("localhost")
	db.Vehicles = db.Session.DB("vehicle_finder").C("vehicles")
	return &db
}

func (db *DB) Close() {
	db.Session.Close()
}
