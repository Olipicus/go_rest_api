package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"code.olipicus.com/go_rest_api/api/message"
	"code.olipicus.com/go_rest_api/api/utility/mongo"
	"github.com/gorilla/mux"
)

// REST Model
type REST struct {
	Collection string
	OBJ        interface{}
}

//UpdateByID ...
func (rest REST) UpdateByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&rest.OBJ)

	if err != nil {
		message.PrintJSONMessage(res, "Format Not Match", "FormatError")
		log.Println(err)
		return
	}
	vars := mux.Vars(req)
	id := vars["id"]

	err = mgh.UpdateData(rest.Collection, id, &rest.OBJ)

	if err != nil {
		message.PrintJSONMessage(res, "Update Fail", "UpdateError")
		log.Println(err)
		return
	}

	message.PrintJSONMessage(res, "Success", "Infomation")
}

//GetDataByID ...
func (rest REST) GetDataByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id := vars["id"]
	obj := mgh.GetOneData(rest.Collection, string(id))

	json.NewEncoder(res).Encode(obj)
}

//InsertData ...
func (rest REST) InsertData(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&rest.OBJ)

	if err != nil {
		message.PrintJSONMessage(res, "Format Not Match", "FormatError")
		log.Println(err)
		return
	}

	err = mgh.InsertData(rest.Collection, &rest.OBJ)

	if err != nil {
		message.PrintJSONMessage(res, "Insert Fail", "UpdateError")
		log.Println(err)
		return
	}

	message.PrintJSONMessage(res, "Success", "Infomation")

}
