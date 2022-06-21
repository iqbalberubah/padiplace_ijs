package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"
)

func HistoryPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var po []e.PO
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	d.DB.Table("po").Where("id_kios = ?", body["id_kios"]).Find(&po)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", po})
}

func DetailPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var po e.PO
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	if result := d.DB.Table("po").Where("id_po = ?", body["id_po"]).First(&po); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", po})
	}
}

func CreatePO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var po e.PO
	json.NewDecoder(r.Body).Decode(&po)
	d.DB.Table("po").Create(&po)
	json.NewEncoder(w).Encode(po)
}
