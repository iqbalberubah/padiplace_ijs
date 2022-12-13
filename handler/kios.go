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
	var user e.User
	var body map[string]interface{}
	buffer, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(buffer, &body); err != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
		return
	}
	if result := d.DB.Table("user").Where("username = ?", body["username"]).First(&user); result.Error != nil {
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
	var kios e.User
	json.NewDecoder(r.Body).Decode(&kios)
	if result := d.DB.Table("user").Where("tlp_kios = ?", kios.TlpUser).First(&kios); result.Error != nil {
		d.DB.Table("user").Create(&kios)
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", kios})
	} else {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Nomor hp sudah terdaftar"})
	}
}

func UpdateTokenFcm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user e.User
	d.DB.Table("user").Where("tlp_user = ?", params["id"]).First(&user)
	json.NewDecoder(r.Body).Decode(&user)
	d.DB.Table("user").Where("tlp_user = ?", params["id"]).Save(&user)
	json.NewEncoder(w).Encode(user)
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
