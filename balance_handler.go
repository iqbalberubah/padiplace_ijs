package main

import (
	"encoding/json"
	"net/http"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Balance
	json.NewDecoder(r.Body).Decode(&user)
	if result := DB.Table("balance").Where("id_peternak = ?", user.IdPeternak).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(ErrorBalance{404, "User tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(SuccesBalance{0, "Succes", user})
	}
}
