package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"code.olipicus.com/go_rest_api/api/person"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)

	//REST API For Person
	router.HandleFunc("/person/{id}", person.GetDataByID).Methods("GET")
	router.HandleFunc("/person", person.InsertData).Methods("POST")
	router.HandleFunc("/person/{id}", person.UpdateByID).Methods("PUT")

	log.Println("Server Start ...")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
