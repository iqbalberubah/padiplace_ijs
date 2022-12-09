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
	d.DB.Table("tbl_product").Find(&products)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", products})
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product e.Product
	if result := d.DB.Table("tbl_product").First(&product, params["id"]); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Product tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", product})
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product e.Product
	json.NewDecoder(r.Body).Decode(&product)
	d.DB.Table("tbl_product").Create(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product e.Product
	d.DB.Table("tbl_product").First(&product, params["id"])
	json.NewDecoder(r.Body).Decode(&product)
	d.DB.Table("tbl_product").Where("id_product = ?", params["id"]).Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product e.Product
	d.DB.Table("tbl_product").Delete(&product, params["id"])
	json.NewEncoder(w).Encode(product)
}
