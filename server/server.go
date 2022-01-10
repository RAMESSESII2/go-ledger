package server

import (
	"log"
	"net/http"

	"github.com/RAMESSESII2/go-ledger/server/repositories"
	"github.com/RAMESSESII2/go-ledger/server/services"
	"github.com/gorilla/mux"
)

var R *mux.Router

func InitializeRouter(address string) {
	R = mux.NewRouter()

	R.HandleFunc("/hello", services.SayHello).Methods("GET")
	R.HandleFunc("/ledger", services.GetLedger).Methods("GET")
	R.HandleFunc("/ledger/{id}", services.GetTransaction).Methods("GET")
	R.HandleFunc("/ledger", services.NewTransaction).Methods("POST")
	R.HandleFunc("/ledger/{id}", services.UpdateTransaction).Methods("PATCH")
	R.HandleFunc("/ledger/{id}", services.DeleteTransaction).Methods("DELETE")

	err := http.ListenAndServe(address, R)
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer(DNS string, address string) {
	repositories.InitialMigration(DNS)
	InitializeRouter(address)
}
