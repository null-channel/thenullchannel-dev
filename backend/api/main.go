package main

import (
	"api/handlers"
	"api/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/s1ntaxe770r/requel"
	"github.com/sirupsen/logrus"
)

var (
	admin_user = os.Getenv("ADMIN_USER")
	admin_pass = os.Getenv("ADMIN_PASS")
)

func main() {
	r := mux.NewRouter()
	r.Use(requel.LogReq)
	r.HandleFunc("/ideas", handlers.GetIdeas).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/ideas", middleware.BasicAuth(handlers.DeleteIdea, admin_user, admin_pass, "")).Methods(http.MethodDelete)
	r.HandleFunc("/vote", handlers.IncrementVote).Methods("POST")
	r.HandleFunc("/create", middleware.BasicAuth(handlers.InsertIdea, admin_user, admin_pass, "")).Methods("POST")
	logrus.Info("server started on 8080")
	handler := cors.Default().Handler(r)
	logrus.Fatal(http.ListenAndServe(":8080", handler))

}
