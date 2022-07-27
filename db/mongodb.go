package db

import (
	"gopkg.in/mgo.v2"
)
var  MongoClient *mgo.Database

func getDB() *mgo.Database {
	if MongoClient !=nil {
		return MongoClient
	}
	session, err := mgo.Dial("mongodb://119.3.85.40:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	MongoClient = session.DB("stock")
	return MongoClient
}
