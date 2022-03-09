package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user e.Peternak
	params := mux.Vars(r)
	if result := d.DB.Table("peternak").Where("tlp_peternak = ? AND password1_user = ?", params["tlp_peternak"], params["password1_user"]).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "User tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", user})
	}
}
