package person

import (
	"encoding/json"
	"log"
	"net/http"

	"code.olipicus.com/go_rest_api/api/message"
	"code.olipicus.com/go_rest_api/api/utility/mongo"
	"github.com/gorilla/mux"
)

const (
	collection = "person"
)

var obj Person

//UpdateByID ...
func UpdateByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&obj)

	if err != nil {
		message.PrintJSONMessage(res, "Format Not Match", "FormatError")
		log.Fatal(err)
		return
	}
	vars := mux.Vars(req)
	id := vars["id"]

	err = mgh.UpdateData("person", id, &obj)

	if err != nil {
		message.PrintJSONMessage(res, "Update Fail", "UpdateError")
		return
	}

	message.PrintJSONMessage(res, "Success", "Infomation")
}

//GetDataByID ...
func GetDataByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id := vars["id"]
	obj := mgh.GetOneData(collection, string(id))

	json.NewEncoder(res).Encode(obj)
}

//InsertData ...
func InsertData(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&obj)

	if err != nil {
		message.PrintJSONMessage(res, "Format Not Match", "FormatError")
		log.Fatal(err)
		return
	}

	err = mgh.InsertData(collection, &obj)

	if err != nil {
		message.PrintJSONMessage(res, "Insert Fail", "UpdateError")
		log.Fatal(err)
		return
	}

	message.PrintJSONMessage(res, "Success", "Infomation")

}
