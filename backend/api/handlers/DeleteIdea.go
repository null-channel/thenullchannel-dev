package handlers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"net/http"
)

func DeleteIdea(resp http.ResponseWriter, req *http.Request) {
	idea := models.Idea{}
	json.NewDecoder(req.Body).Decode(&idea)
	dbcon := db.Connect()
	err := db.DeleteIdea(dbcon, idea)
	if err != nil {
		http.Error(resp, err.Error(), 405)
		return
	}
	resp.WriteHeader(http.StatusOK)
}
