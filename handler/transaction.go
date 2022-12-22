package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func TransactionHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction []e.TransactionHistory
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	d.DB.Table("tbl_transaksi").Select("tbl_transaksi.id_transaksi, tbl_transaksi.date, tbl_transaksi.keterangan, tbl_transaksi.balance, tbl_transaksi.id_product, tbl_transaksi.id_user, tbl_transaksi.type_trk, tbl_transaksi.status_transaksi, tbl_transaksi.number_trx, tbl_transaksi.jumlah_trx, tbl_product.nm_product, tbl_product.foto").Joins("left join tbl_product on tbl_product.id_product = tbl_transaksi.id_product").Where("tbl_transaksi.id_user = ?", body["id_user"]).Find(&transaction)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", transaction})
}

func TransactionDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction e.Transaction
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	if result := d.DB.Table("tbl_transaksi").Where("id_transaksi = ?", body["id_transaksi"]).First(&transaction); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", transaction})
	}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transaction e.Transaction
	json.NewDecoder(r.Body).Decode(&transaction)
	d.DB.Table("tbl_transaksi").Create(&transaction)
	json.NewEncoder(w).Encode(transaction)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var po e.Transaction
	d.DB.Table("tbl_transaksi").First(&po, params["id"])
	json.NewDecoder(r.Body).Decode(&po)
	d.DB.Table("tbl_transaksi").Where("number_trx = ?", params["id"]).Save(&po)
	json.NewEncoder(w).Encode(po)
}
