package main

import (
	"log"
	"net/http"

	"link-converter/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/converter/todeeplink", handlers.ResponseHandler(handlers.ToDeepLink)).Methods("POST")
	router.Handle("/converter/toweburl", handlers.ResponseHandler(handlers.ToWebUrl)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
