package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user e.Peternak
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	if result := d.DB.Table("peternak").Where("tlp_peternak = ? AND password1_user = ?", body["tlp_peternak"], body["password1_user"]).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "User tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", user})
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var peternak e.Peternak
	json.NewDecoder(r.Body).Decode(&peternak)
	if result := d.DB.Table("peternak").Where("tlp_peternak = ?", peternak.TlpPeternak).First(&peternak); result.Error != nil {
		d.DB.Table("peternak").Create(&peternak)
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", peternak})
	} else {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Nomor hp sudah terdaftar"})
	}
}

func UpdateTokenFcm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var peternak e.Peternak
	d.DB.Table("peternak").Where("tlp_peternak = ?", params["id"]).First(&peternak)
	json.NewDecoder(r.Body).Decode(&peternak)
	d.DB.Table("peternak").Where("tlp_peternak = ?", params["id"]).Save(&peternak)
	json.NewEncoder(w).Encode(peternak)
}
