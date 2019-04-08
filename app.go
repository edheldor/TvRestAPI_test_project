package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/tv", a.GetAllTv).Methods("GET")
	a.Router.HandleFunc("/tv/{id}", a.GetTv).Methods("GET")
	a.Router.HandleFunc("/tv", a.CreateTv).Methods("POST")
	a.Router.HandleFunc("/tv/{id}", a.DeleteTv).Methods("DELETE")
	a.Router.HandleFunc("/tv/{id}", a.EditTv).Methods("PUT")

}

func (a *App) GetAllTv(w http.ResponseWriter, r *http.Request) {
	tvs, err := getAllTv(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tvs)
}

func (a *App) CreateTv(w http.ResponseWriter, r *http.Request) {
	var tv tv
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tv); err != nil {
		respondWithError(w, http.StatusBadRequest, "Неправильный формать JSON")
		return
	}
	defer r.Body.Close()

	if err := tv.createTv(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, "Тв добавлен")
}

func (a *App) GetTv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := TvIdChecker(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID должен быть больше нуля и быть числом")
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неправильный  ID")
		return
	}

	tv := tv{Id: id}
	if err := tv.getTV(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "TV с таким ID не найден")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, tv)
}

func (a *App) EditTv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := TvIdChecker(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID должен быть больше нуля и быть числом")
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неправильный  ID")
		return
	}

	var tv tv
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tv); err != nil {
		respondWithError(w, http.StatusBadRequest, "Неправильный формать JSON")
		return
	}
	defer r.Body.Close()
	tv.Id = id

	if err := tv.updateTv(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tv)
}

func (a *App) DeleteTv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := TvIdChecker(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID должен быть больше нуля и быть числом")
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Неправильный  ID")
		return
	}

	tv := tv{Id: id}
	if err := tv.deleteTv(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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

func TvChecker(tv tv) error {
	if len(tv.Manufacturer) < 3 {
		return fmt.Errorf("В поле Manufacturer должно быть 3 и более символов")
	}

	if len(tv.Model) < 2 {
		return fmt.Errorf("В поле Manufacturer должно быть 2 и более символов")
	}

	if tv.Year < 2010 {
		return fmt.Errorf("год должен быть больше 2010")
	}
	return nil

}
