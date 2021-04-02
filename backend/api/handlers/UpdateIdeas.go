package handlers

import (
	"api/db"
	"api/models"
	"encoding/json"
	"net/http"
)

func UpdateIdea(resp http.ResponseWriter, req *http.Request) {
	idea := models.Idea{}
	// Todo: Data validation
	json.NewDecoder(req.Body).Decode(&idea)
	dbcon := db.Connect()
	err := db.UpdateIdea(&idea, dbcon)
	if err != nil {
		http.Error(resp, "unable to update post", 500)
	}
}
