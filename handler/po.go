package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"
)

func GetHistoryPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user e.PO
	var po []e.PO
	json.NewDecoder(r.Body).Decode(&user)
	if result := d.DB.Table("po").Where("id_peternak = ?", user.IdPeternak).Find(&po); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", po})
	}
}

func GetPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user e.PO
	json.NewDecoder(r.Body).Decode(&user)
	if result := d.DB.Table("po").Where("id_po = ?", user.IdPo).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", user})
	}
}
