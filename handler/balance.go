package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user e.Balance
	json.NewDecoder(r.Body).Decode(&user)
	if result := d.DB.Table("balance").Where("id_peternak = ?", user.IdPeternak).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorBalance{404, "User tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesBalance{0, "Succes", user})
	}
}
