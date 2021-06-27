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
	parserr := json.NewDecoder(req.Body).Decode(&idea)
	if parserr != nil {
		http.Error(resp,"could not parse request body:"+parserr.Error(),http.StatusBadRequest)
		return
	}
	err := db.UpdateIdea(&idea)
	if err != nil {
		http.Error(resp, "unable to update post", 500)
	}
}
