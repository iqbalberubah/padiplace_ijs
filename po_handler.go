package main

import (
	"encoding/json"
	"net/http"
)

func GetHistoryPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user PO
	var po []PO
	json.NewDecoder(r.Body).Decode(&user)
	if result := DB.Table("po").Where("id_peternak = ?", user.IdPeternak).Find(&po); result.Error != nil {
		json.NewEncoder(w).Encode(ErrorPO{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(SuccesPOData{0, "Succes", po})
	}
}

func GetPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user PO
	json.NewDecoder(r.Body).Decode(&user)
	if result := DB.Table("po").Where("id_po = ?", user.IdPo).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(ErrorPO{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(SuccesPO{0, "Succes", user})
	}
}
