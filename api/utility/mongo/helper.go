package mongo

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName       = "test"
	mongoAddress = "192.168.99.100:27017"
)

//Helper Struct of MongoHelper
type Helper struct {
	session *mgo.Session
}

//Init : Initial DB
func (h *Helper) Init() {
	session, err := mgo.Dial(mongoAddress)
	if err != nil {
		log.Fatal(err)
	}
	h.session = session
	log.Println("Connect MongoDB OK!")

	// Optional. Switch the session to a monotonic behavior.
	h.session.SetMode(mgo.Monotonic, true)
}

//Close : Close DB Session
func (h *Helper) Close() {
	h.session.Close()
}

//GetOneData : Get Single Document
func (h *Helper) GetOneData(collectionName string, id string) interface{} {
	c := h.session.DB(dbName).C(collectionName)
	var obj interface{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&obj)

	if err != nil {
		return nil
	}

	return obj
}

//InsertData : Insert Document to Collection
func (h *Helper) InsertData(collectionName string, obj interface{}) error {
	c := h.session.DB(dbName).C(collectionName)
	return c.Insert(obj)
}

//UpdateData : Update Document
func (h *Helper) UpdateData(collectionName string, id string, obj interface{}) error {
	c := h.session.DB(dbName).C(collectionName)
	update := bson.M{"$set": obj}
	if err := c.UpdateId(bson.ObjectIdHex(id), update); err != nil {
		return err
	}
	return nil
}
