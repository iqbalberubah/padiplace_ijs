package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var balance e.Balance
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	if result := d.DB.Table("tbl_balance").Where("id_user = ?", body["id_user"]).First(&balance); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Data tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", balance})
	}
}
