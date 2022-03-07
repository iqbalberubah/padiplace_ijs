package main

import (
	"encoding/json"
	"net/http"
)

func GetPO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user PO
	json.NewDecoder(r.Body).Decode(&user)
	if result := DB.Table("po").Where("id_peternak = ?").First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(ErrorPO{404, "PO tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(SuccesPO{0, "Succes", user})
	}
}
