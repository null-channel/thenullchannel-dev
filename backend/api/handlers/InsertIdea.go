package handlers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"net/http"
)

func InsertIdea(resp http.ResponseWriter, req *http.Request) {
	idea := models.Idea{}
	// Todo: Data validation
	json.NewDecoder(req.Body).Decode(&idea)
	dbcon := db.Connect()
	result, err := db.Insert(idea, dbcon)
	if err != nil {
		http.Error(resp, "could not add idea to the database. "+err.Error(), http.StatusBadRequest)
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
