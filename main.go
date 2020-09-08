package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	connectToDatabase()
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
