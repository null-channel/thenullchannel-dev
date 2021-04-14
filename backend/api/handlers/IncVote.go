package handlers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"net/http"
)

func IncrementVote(resp http.ResponseWriter, req *http.Request) {
	idea := models.Idea{}
	json.NewDecoder(req.Body).Decode(&idea)
	dbcon := db.Connect()
	queryerr := db.IncVote(idea.ID, dbcon)
	if queryerr != nil {
		http.Error(resp, "could not increase vote with id", 500)
		return
	}
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Content-Type", "application/json")

	if req.Method == http.MethodOptions {
		return
	}

	resp.WriteHeader(http.StatusOK)
}
