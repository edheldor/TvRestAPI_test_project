package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tv", GetAllTv).Methods("GET")
	router.HandleFunc("/tv/{id}", GetTv).Methods("GET")
	router.HandleFunc("/tv/{id}", CreateTv).Methods("POST")
	router.HandleFunc("/tv/{id}", DeleteTv).Methods("DELETE")
	router.HandleFunc("/tv/{id}", EditTv).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetAllTv(w http.ResponseWriter, r *http.Request) {

}

func GetTv(w http.ResponseWriter, r *http.Request) {

}

func CreateTv(w http.ResponseWriter, r *http.Request) {

}

func DeleteTvTv(w http.ResponseWriter, r *http.Request) {

}

func EditTvTv(w http.ResponseWriter, r *http.Request) {

}
