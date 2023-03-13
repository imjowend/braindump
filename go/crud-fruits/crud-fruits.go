package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Fruit struct {
	ID             string    `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Quantity       int       `json:"quantity,omitempty"`
	Price          float64   `json:"price,omitempty"`
	DateCreated    time.Time `json:"date_created,omitempty"`
	DateLastUpdate time.Time `json:"date_last_updated,omitempty"`
	Owner          string    `json:"owner,omitempty"`
	Status         string    `json:"status,omitempty"`
}

var fruits []Fruit

func GetFruitByIDEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range fruits {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Fruit{})
}

func GetFruitsEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(fruits)
}

func CreateFruitEndpoint(w http.ResponseWriter, req *http.Request) {
	var fruit Fruit
	_ = json.NewDecoder(req.Body).Decode(&fruit)
	fruit.ID = uuid.New().String()
	fruit.DateCreated = time.Now()
	fruit.DateLastUpdate = time.Now()
	fruits = append(fruits, fruit)
	json.NewEncoder(w).Encode(fruit)
}

func UpdateFruitEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range fruits {
		if item.ID == params["id"] {
			var fruit Fruit
			_ = json.NewDecoder(req.Body).Decode(&fruit)
			fruit.ID = item.ID
			fruit.DateCreated = item.DateCreated
			fruit.DateLastUpdate = time.Now()
			fruits[index] = fruit
			json.NewEncoder(w).Encode(fruit)
			return
		}
	}
	json.NewEncoder(w).Encode(fruits)
}

func DeleteFruitEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range fruits {
		if item.ID == params["id"] {
			fruits = append(fruits[:index], fruits[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(fruits)
}

func main() {
	router := mux.NewRouter()

	fruits = append(fruits, Fruit{ID: "1", Name: "Manzana", Quantity: 12, Price: 1000, DateCreated: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), DateLastUpdate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), Owner: "test", Status: "comestible"})

	router.HandleFunc("/fruits", GetFruitsEndpoint).Methods("GET")
	router.HandleFunc("/fruits/{id}", GetFruitByIDEndpoint).Methods("GET")
	router.HandleFunc("/fruits", CreateFruitEndpoint).Methods("POST")
	router.HandleFunc("/fruits/{id}", UpdateFruitEndpoint).Methods("PUT")
	router.HandleFunc("/fruits/{id}", DeleteFruitEndpoint).Methods("DELETE")

	fmt.Println("a")
}
