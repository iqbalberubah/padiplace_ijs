package entity

type Transaction struct {
	// gorm.Model
	IdTransaksi     int    `json:"id_transaksi"`
	Date            string `json:"date"`
	Keterangan      string `json:"keterangan"`
	Balance         string `json:"balance"`
	NumberTrx       string `json:"number_trx"`
	IdProduct       int    `json:"id_product"`
	IdUser          int    `json:"id_user"`
	TypeTrk         int    `json:"type_trk"`
	StatusTransaksi int    `json:"status_transaksi"`
	JumlahTrx       string `json:"jumlah_trx"`
}

type TransactionHistory struct {
	// gorm.Model
	IdTransaksi     int    `json:"id_transaksi"`
	Date            string `json:"date"`
	Keterangan      string `json:"keterangan"`
	Balance         string `json:"balance"`
	NumberTrx       string `json:"number_trx"`
	IdProduct       int    `json:"id_product"`
	IdUser          int    `json:"id_user"`
	TypeTrk         int    `json:"type_trk"`
	StatusTransaksi int    `json:"status_transaksi"`
	NmProduct       string `json:"nm_product"`
	JumlahTrx       string `json:"jumlah_trx"`
}
