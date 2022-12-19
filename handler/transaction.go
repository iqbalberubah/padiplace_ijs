package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func HistoryPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var po []e.Transaction
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	d.DB.Table("tbl_transaksi").Where("id_user = ?", body["id_user"]).Find(&po)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", po})
}

func DetailPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var po e.Transaction
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	if result := d.DB.Table("tbl_transaksi").Where("number_trx = ?", body["number_trx"]).First(&po); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", po})
	}
}

func CreatePO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var po e.Transaction
	json.NewDecoder(r.Body).Decode(&po)
	d.DB.Table("tbl_transaksi").Create(&po)
	json.NewEncoder(w).Encode(po)
}

func UpdatePO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var po e.Transaction
	d.DB.Table("tbl_transaksi").First(&po, params["id"])
	json.NewDecoder(r.Body).Decode(&po)
	d.DB.Table("tbl_transaksi").Where("number_trx = ?", params["id"]).Save(&po)
	json.NewEncoder(w).Encode(po)
}
