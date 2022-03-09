package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []e.Product
	d.DB.Table("product").Find(&products)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", products})
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var products e.Product
	if result := d.DB.Table("product").First(&products, params["id"]); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Product tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", products})
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products e.Product
	json.NewDecoder(r.Body).Decode(&products)
	d.DB.Table("product").Create(&products)
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var products e.Product
	d.DB.Table("product").First(&products, params["id"])
	json.NewDecoder(r.Body).Decode(&products)
	d.DB.Table("product").Save(&products)
	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var products e.Product
	d.DB.Table("product").Delete(&products, params["id"])
	json.NewEncoder(w).Encode(products)
}
