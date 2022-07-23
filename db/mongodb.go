package db

import (
	"gopkg.in/mgo.v2"
)


func getDB() *mgo.Database {
	session, err := mgo.Dial("mongodb://119.3.85.40:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("stock")
	return db
}
