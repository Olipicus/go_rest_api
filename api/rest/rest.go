package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"code.olipicus.com/go_rest_api/api/message"
	"code.olipicus.com/go_rest_api/api/utility/mongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// REST Model
type REST struct {
	Collection string
	OBJ        interface{}
}

//RemoveByID ...
func (rest *REST) RemoveByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	vars := mux.Vars(req)
	id := vars["id"]

	msg := "Success"
	msgType := "Information"

	if err := mgh.RemoveByID(rest.Collection, id); err != nil {
		switch err.Error() {
		case "not found":
			msg, msgType = "Data Not Found", "DeleteFail"
		default:
			msg, msgType = "Delete Fail", "DeleteFail"
		}
		log.Println(err)
	}
	message.PrintJSONMessage(res, msg, msgType)
}

//UpdateByID ...
func (rest *REST) UpdateByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	decoder := json.NewDecoder(req.Body)

	msg := "Success"
	msgType := "Information"

	if err := decoder.Decode(&rest.OBJ); err != nil {
		msg, msgType = "Format Not Match", "FormatError"
		log.Println(err)
	} else {
		vars := mux.Vars(req)
		id := vars["id"]

		if err = mgh.UpdateData(rest.Collection, id, &rest.OBJ); err != nil {
			switch err {
			case mgo.ErrNotFound:
				msg, msgType = mgo.ErrNotFound.Error(), "UpdateFail"
			default:
				msg, msgType = "Update Fail", "UpdateFail"
			}
			log.Println(err)
		}
	}

	message.PrintJSONMessage(res, msg, msgType)
}

//GetDataByID ...
func (rest *REST) GetDataByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	id := vars["id"]

	var obj interface{}
	var err error
	if obj, err = mgh.GetOneData(rest.Collection, string(id)); err != nil {
		var msg, msgType string
		switch err {
		case mgo.ErrNotFound:
			msg, msgType = mgo.ErrNotFound.Error(), "UpdateFial"
		default:
			msg, msgType = "Get Data Fail", "Error"
		}
		message.PrintJSONMessage(res, msg, msgType)
	} else {
		json.NewEncoder(res).Encode(obj)
	}

}

//InsertData ...
func (rest *REST) InsertData(res http.ResponseWriter, req *http.Request) {
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
