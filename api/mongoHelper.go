package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "test"
const mongoAddress = "192.168.99.100:27017"

func getOneData(collectionName string, id string) interface{} {
	session, err := mgo.Dial(mongoAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connect MongoDB OK!")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	c := session.DB(dbName).C(collectionName)
	var obj interface{}
	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&obj)

	if err != nil {
		return nil
	}

	return obj
}

func insertData(collectionName string, obj interface{}) error {
	session, err := mgo.Dial(mongoAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connect MongoDB OK!")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	c := session.DB(dbName).C(collectionName)
	err = c.Insert(obj)

	return err

}
