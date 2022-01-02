package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/ledger", GetLedger).Methods("GET")
	r.HandleFunc("/ledger/{id}", GetTransaction).Methods("GET")
	r.HandleFunc("/ledger", NewTransaction).Methods("POST")
	r.HandleFunc("/ledger/{id}", UpdateTransaction).Methods("PUT")
	r.HandleFunc("/ledger/{id}", DeleteTransaction).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func startServer() {
	InitialMigration()
	initializeRouter()
}
