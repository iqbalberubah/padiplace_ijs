package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []Product
	DB.Table("product").Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user Product
	DB.Table("product").First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Product
	json.NewDecoder(r.Body).Decode(&user)
	DB.Table("product").Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user Product
	DB.Table("product").First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Table("product").Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user Product
	DB.Table("product").Delete(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
