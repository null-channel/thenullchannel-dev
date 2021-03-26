package main

import (
	"api/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/s1ntaxe770r/requel"
	"github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	r.Use(requel.LogReq)
	r.HandleFunc("/ideas", handlers.GetIdeas).Methods("GET")
	r.HandleFunc("/vote", handlers.IncrementVote).Methods("POST")

	logrus.Info("server started on 8080")
	logrus.Fatal(http.ListenAndServe(":8080", r))
}
