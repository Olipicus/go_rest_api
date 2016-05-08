package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"code.olipicus.com/go_rest_api/api/utility/mongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// REST Model
type REST struct {
	Collection string
	OBJ        interface{}
}

func (rest *REST) responseResult(res http.ResponseWriter, result Result) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(result.StatusCode)
	json.NewEncoder(res).Encode(result)
}

func (rest *REST) responseDataResult(res http.ResponseWriter, result Result, data interface{}) {
	resultSuccess := ResultSuccess{
		Result: result,
		Data:   data,
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resultSuccess)
}

func (rest *REST) responseErrorResult(res http.ResponseWriter, err Error) {
	result := ResultError{
		Result: Result{
			StatusCode:  http.StatusInternalServerError,
			Description: "Error",
		},
		Error: err,
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(res).Encode(result)
}

//RemoveByID : Remove data by HTTP DELETE
func (rest *REST) RemoveByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	vars := mux.Vars(req)
	id := vars["id"]

	if err := mgh.RemoveByID(rest.Collection, id); err != nil {
		log.Println(err)
		switch err {
		case mgo.ErrNotFound:
			rest.responseResult(res, resultDataNotFound)
		default:
			rest.responseErrorResult(res, Error{Code: 500, ErrorMessage: err.Error()})
		}
		return
	}
	rest.responseResult(res, resultDeleteComplete)
}

//UpdateByID : Update Data by HTTP PUT
func (rest *REST) UpdateByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&rest.OBJ); err != nil {
		log.Println(err)
		rest.responseErrorResult(res, Error{Code: 500, ErrorMessage: err.Error()})
		return
	}
	vars := mux.Vars(req)
	id := vars["id"]

	if err := mgh.UpdateData(rest.Collection, id, &rest.OBJ); err != nil {
		log.Println(err)
		switch err {
		case mgo.ErrNotFound:
			rest.responseResult(res, resultDataNotFound)
		default:
			rest.responseErrorResult(res, Error{Code: 500, ErrorMessage: err.Error()})
		}
		return
	}

	rest.responseResult(res, resultUpdateComplete)
}

//GetDataByID : Get Single Data By HTTP GET
func (rest *REST) GetDataByID(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	vars := mux.Vars(req)
	id := vars["id"]

	var obj interface{}
	var err error
	if obj, err = mgh.GetOneData(rest.Collection, string(id)); err != nil {
		log.Println(err)
		switch err {
		case mgo.ErrNotFound:
			rest.responseResult(res, resultDataNotFound)
		default:
			rest.responseErrorResult(res, Error{Code: 500, ErrorMessage: err.Error()})
		}
		return
	}
	rest.responseDataResult(res, resultSuccess, obj)
}

//InsertData : Insert Data By HTTP POST
func (rest *REST) InsertData(res http.ResponseWriter, req *http.Request) {
	var mgh mongo.Helper
	mgh.Init()

	defer mgh.Close()

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&rest.OBJ); err != nil {
		log.Println(err)
		rest.responseErrorResult(res, Error{Code: 500, ErrorMessage: err.Error()})
		return
	}

	if err := mgh.InsertData(rest.Collection, &rest.OBJ); err != nil {
		log.Println(err)
		rest.responseErrorResult(res, Error{Code: 500, ErrorMessage: err.Error()})
		return
	}

	rest.responseResult(res, resultInsertComplete)
}
