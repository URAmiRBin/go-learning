package main

import (
	"log"
	"net/http"

	"github.com/URAmiRBin/go-learning/login/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	r.HandleFunc("/profile", handler.ProfileHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
