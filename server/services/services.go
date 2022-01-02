package services

import (
	"encoding/json"
	"net/http"

	"github.com/RAMESSESII2/go-ledger/server/models"
	"github.com/RAMESSESII2/go-ledger/server/repositories"
	"github.com/gorilla/mux"
)

func GetLedger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "hi!")
	var ledger []models.Transaction
	repositories.DB.Find(&ledger)
	json.NewEncoder(w).Encode(ledger)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction models.Transaction
	repositories.DB.First(&trasaction, params["id"])
	json.NewEncoder(w).Encode(trasaction)
}
func NewTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var trasaction models.Transaction
	json.NewDecoder(r.Body).Decode(&trasaction)
	repositories.DB.Create(&trasaction)
	json.NewEncoder(w).Encode(trasaction)
}
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction models.Transaction
	repositories.DB.First(&trasaction, params["id"])
	json.NewDecoder(r.Body).Decode(&trasaction)
	repositories.DB.Save(&trasaction)
	json.NewEncoder(w).Encode(trasaction)
}
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var trasaction models.Transaction
	repositories.DB.Delete(&trasaction, params["id"])
	json.NewEncoder(w).Encode("The transaction is deleted successfully!")
}
