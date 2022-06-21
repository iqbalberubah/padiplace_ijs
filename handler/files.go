package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func GetFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var files []e.Files
	d.DB.Table("file").Find(&files)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", files})
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var files e.Files
	if result := d.DB.Table("file").First(&files, params["id"]); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "File tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", files})
	}
}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var files e.Files
	json.NewDecoder(r.Body).Decode(&files)
	d.DB.Table("file").Create(&files)
	json.NewEncoder(w).Encode(files)
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var files e.Files
	d.DB.Table("file").First(&files, params["id"])
	json.NewDecoder(r.Body).Decode(&files)
	d.DB.Table("file").Where("id_file = ?", params["id"]).Save(&files)
	json.NewEncoder(w).Encode(files)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var files e.Files
	d.DB.Table("file").Delete(&files, params["id"])
	json.NewEncoder(w).Encode(files)
}
