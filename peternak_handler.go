package main

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user Peternak
	json.NewDecoder(r.Body).Decode(&user)
	if result := DB.Table("peternak").Where("tlp_peternak = ? AND password1_user = ?", user.TlpPeternak, user.Password1User).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(ErrorPeternak{404, "User tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(SuccesPeternak{0, "Succes", user})
	}
}
