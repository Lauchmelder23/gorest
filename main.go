package main

import (
	"fmt";
	"log";
	"net/http";
	"github.com/gorilla/mux";
)

func main() {
	router := registerEndpoints()

	fmt.Println("Starting server at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func registerEndpoints() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", ShowMessage).Methods("GET")
	router.HandleFunc("/", SaveValue).Methods("POST")

	return router
}
