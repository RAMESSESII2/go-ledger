package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "hi!")
	var ledger []Transaction
	DB.Find(&ledger)
	json.NewEncoder(w).Encode(ledger)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction Transaction
	DB.First(&trasaction, params["id"])
	json.NewEncoder(w).Encode(trasaction)
}
func NewTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var trasaction Transaction
	json.NewDecoder(r.Body).Decode(&trasaction)
	DB.Create(&trasaction)
	json.NewEncoder(w).Encode(trasaction)
}
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction Transaction
	DB.First(&trasaction, params["id"])
	json.NewDecoder(r.Body).Decode(&trasaction)
	DB.Save(&trasaction)
	json.NewEncoder(w).Encode(trasaction)
}
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction Transaction
	DB.Delete(&trasaction, params["id"])
	json.NewEncoder(w).Encode("The transaction is deleted successfully!")
}
