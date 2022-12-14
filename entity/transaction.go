package entity

type Transaction struct {
	// gorm.Model
	IdTransaksi     int    `json:"id_transaksi"`
	Date            string `json:"date"`
	Keterangan      string `json:"keterangan"`
	Balance         string `json:"balance"`
	IdProduct       int    `json:"id_product"`
	IdUser          int    `json:"id_user"`
	TypeTrk         int    `json:"type_trk"`
	StatusTransaksi int    `json:"status_transaksi"`
}
