package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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
	if result := d.DB.Table("peternak").Where("tlp_peternak = ?", body["tlp_peternak"]).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "User tidak ditemukan"})
	} else {
		match := PasswordVerify(body["password1_user"].(string), user.Password1User)
		if match {
			json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", user})
		} else {
			json.NewEncoder(w).Encode(e.ErrorResponse{404, "Password tidak sesuai"})
		}
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

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
