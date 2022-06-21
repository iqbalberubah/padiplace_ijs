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
	var user e.Kios
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	if result := d.DB.Table("kios").Where("tlp_kios = ?", body["tlp_kios"]).First(&user); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "User tidak ditemukan"})
	} else {
		match := PasswordVerify(body["password"].(string), user.Password)
		if match {
			json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", user})
		} else {
			json.NewEncoder(w).Encode(e.ErrorResponse{404, "Password tidak sesuai"})
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var kios e.Kios
	json.NewDecoder(r.Body).Decode(&kios)
	if result := d.DB.Table("kios").Where("tlp_kios = ?", kios.TlpKios).First(&kios); result.Error != nil {
		d.DB.Table("kios").Create(&kios)
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", kios})
	} else {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Nomor hp sudah terdaftar"})
	}
}

func UpdateTokenFcm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var kios e.Kios
	d.DB.Table("kios").Where("tlp_kios = ?", params["id"]).First(&kios)
	json.NewDecoder(r.Body).Decode(&kios)
	d.DB.Table("kios").Where("tlp_kios = ?", params["id"]).Save(&kios)
	json.NewEncoder(w).Encode(kios)
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
