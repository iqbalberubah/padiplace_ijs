package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []e.ProductList
	d.DB.Table("tbl_product").Select("tbl_product.id_product, tbl_product.nm_product, tbl_product.harga_product, tbl_product.stock, tbl_product.type_product, tbl_product.satuan_product, tbl_product.des_product, tbl_product.foto, tbl_product.foto2, tbl_product.foto3, tbl_product.foto4, tbl_product.foto5, tbl_product.foto6, tbl_product.status, tbl_product.id_user, tbl_product.tgl_date, tbl_product.kom_product, tbl_product.tips_peng_product, tbl_product.no_sku, tbl_product.publish_s, user.nama_prusahaan").Joins("left join user on user.id_user = tbl_product.id_user").Find(&products)
	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", products})

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product e.ProductList
	if result := d.DB.Table("tbl_product").Select("tbl_product.id_product, tbl_product.nm_product, tbl_product.harga_product, tbl_product.stock, tbl_product.type_product, tbl_product.satuan_product, tbl_product.des_product, tbl_product.foto, tbl_product.foto2, tbl_product.foto3, tbl_product.foto4, tbl_product.foto5, tbl_product.foto6, tbl_product.status, tbl_product.id_user, tbl_product.tgl_date, tbl_product.kom_product, tbl_product.tips_peng_product, tbl_product.no_sku, tbl_product.publish_s, user.nama_prusahaan").Joins("left join user on user.id_user = tbl_product.id_user").First(&product, params["id"]); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Product tidak ditemukan"})
	} else {
		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", product})
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product e.Product
	json.NewDecoder(r.Body).Decode(&product)
	d.DB.Table("tbl_product").Create(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product e.Product
	d.DB.Table("tbl_product").First(&product, params["id"])
	json.NewDecoder(r.Body).Decode(&product)
	d.DB.Table("tbl_product").Where("id_product = ?", params["id"]).Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product e.Product
	d.DB.Table("tbl_product").Delete(&product, params["id"])
	json.NewEncoder(w).Encode(product)
}
