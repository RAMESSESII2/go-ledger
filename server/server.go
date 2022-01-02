package server

import (
	"log"
	"net/http"

	"github.com/RAMESSESII2/go-ledger/server/repositories"
	"github.com/RAMESSESII2/go-ledger/server/services"
	"github.com/gorilla/mux"
)

func InitializeRouter(address string) {
	r := mux.NewRouter()

	r.HandleFunc("/ledger", services.GetLedger).Methods("GET")
	r.HandleFunc("/ledger/{id}", services.GetTransaction).Methods("GET")
	r.HandleFunc("/ledger", services.NewTransaction).Methods("POST")
	r.HandleFunc("/ledger/{id}", services.UpdateTransaction).Methods("PATCH")
	r.HandleFunc("/ledger/{id}", services.DeleteTransaction).Methods("DELETE")

	err := http.ListenAndServe(address, r)
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer(DNS string, address string) {
	repositories.InitialMigration(DNS)
	InitializeRouter(address)
}
