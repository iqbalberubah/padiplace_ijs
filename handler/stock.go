package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func GetStocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stocks []e.Stock
	d.DB.Table("stock").Find(&stocks)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", stocks})
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var stock e.Stock
	if result := d.DB.Table("stock").First(&stock, params["id"]); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Product tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", stock})
	}
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stock e.Stock
	json.NewDecoder(r.Body).Decode(&stock)
	d.DB.Table("stock").Create(&stock)
	json.NewEncoder(w).Encode(stock)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var stock e.Stock
	d.DB.Table("stock").First(&stock, params["id"])
	json.NewDecoder(r.Body).Decode(&stock)
	d.DB.Table("stock").Where("id_stock = ?", params["id"]).Save(&stock)
	json.NewEncoder(w).Encode(stock)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var stock e.Stock
	d.DB.Table("stock").Delete(&stock, params["id"])
	json.NewEncoder(w).Encode(stock)
}
