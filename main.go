package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Item struct {
	ID		int `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type HealthResponse struct {
	IsAlive bool `json:"alive"`
}

func main() {
	fmt.Printf("henkin4.0 starting")
	router := mux.NewRouter()
	router.HandleFunc("/", GetItem).Methods("GET")
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { OK(w, HealthResponse{IsAlive:true}) })
	//router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	//router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	//router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func OK(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	response := Item{ID: 1, Name: "FirstItem", CreatedAt: time.Now()}
	OK(w, response)
}
