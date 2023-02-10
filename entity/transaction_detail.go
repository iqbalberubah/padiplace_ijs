package entity

type TransactionDetail struct {
	// gorm.Model
	IdDetailtrx int `json:"id_detailtrx"`
	IdTransaksi int `json:"id_transaksi"`
	IdProduct   int `json:"id_product"`
}
