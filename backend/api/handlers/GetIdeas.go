package handlers

import (
	"api/db"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetIdeas(resp http.ResponseWriter, req *http.Request) {
	dbcon := db.Connect()
	ideas := db.GetIdeas(dbcon)
	jsonresp, err := json.Marshal(ideas)
	if err != nil {
		logrus.Panicf("unable to marshal json , reason %s ", err.Error())
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonresp)

}
