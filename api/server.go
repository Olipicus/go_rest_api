package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"code.olipicus.com/go_rest_api/api/model"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/person/{personID}", retrievePerson).Methods("GET")
	router.HandleFunc("/person", createPerson).Methods("POST")

	log.Println("Server Start ...")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func retrievePerson(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	personID := vars["personID"]
	person := getOneData("person", string(personID))

	if person != nil {
		log.Println("Request Person [" + personID + "] Success!")
	} else {
		log.Println("Request Person [" + personID + "] Not Found!")
	}

	json.NewEncoder(res).Encode(person)
}

func createPerson(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var p model.Person
	err := decoder.Decode(&p)

	var resMessage model.Message

	if err != nil {
		resMessage.Message = "Format Not Match"
		resMessage.MessageType = "Error"
		log.Fatal(err)

	} else {
		err = insertData("person", &p)
		if err != nil {
			resMessage.Message = "Insert Fail"
			resMessage.MessageType = "Error"
		} else {
			resMessage.Message = "Success"
			resMessage.MessageType = "Infomation"
		}
	}

	json.NewEncoder(res).Encode(resMessage)
}
