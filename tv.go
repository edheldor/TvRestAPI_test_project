package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// our main function
var tvs []Tv

func main() {

	//настраиваем роуты для веб-приложения
	router := mux.NewRouter()
	router.HandleFunc("/tv", GetAllTv).Methods("GET")
	router.HandleFunc("/tv/{id}", GetTv).Methods("GET")
	router.HandleFunc("/tv", CreateTv).Methods("POST")
	router.HandleFunc("/tv/{id}", DeleteTv).Methods("DELETE")
	router.HandleFunc("/tv/{id}", EditTv).Methods("PUT")
	//запускаем веб-сервер и добавляем логгирование
	log.Fatal(http.ListenAndServe(":80", router))
}

func GetAllTv(w http.ResponseWriter, r *http.Request) {
	tvid := mux.Vars(r)["id"]
	err := TvIdChecker(tvid)
	if err != nil {
		log.Println(err)

	} else {
		id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		log.Println(id)
	}

}

func GetTv(w http.ResponseWriter, r *http.Request) {
	tvid := mux.Vars(r)["id"]
	err := TvIdChecker(tvid)
	if err != nil {
		log.Println(err)

	} else {
		id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		log.Println(id)
	}

}

func CreateTv(w http.ResponseWriter, r *http.Request) {
	log.Println(r.PostForm)
}

func DeleteTv(w http.ResponseWriter, r *http.Request) {
	tvid := mux.Vars(r)["id"]
	err := TvIdChecker(tvid)
	if err != nil {
		log.Println(err)

	} else {
		id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		log.Println(id)
	}

}

func EditTv(w http.ResponseWriter, r *http.Request) {

}

//функция для проверки введеннго ID телевизора
func TvIdChecker(input_id string) error {

	id, err := strconv.ParseInt(input_id, 10, 64)
	if err == nil {
		if id < 0 {
			return fmt.Errorf("ID должен быть больше нуля")
		} else {
			return nil
		}
	} else {
		return err
	}

}

type Tv struct {
	id           int64
	brand        string
	manufacturer string
	model        string
	year         uint8
}
