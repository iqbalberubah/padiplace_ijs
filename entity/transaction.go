package entity

type Transaction struct {
	// gorm.Model
	IdTransaksi      int    `json:"id_transaksi"`
	Date             string `json:"date"`
	Keterangan       string `json:"keterangan"`
	Balance          string `json:"balance"`
	NumberTrx        string `json:"number_trx"`
	IdProduct        int    `json:"id_product"`
	IdUser           int    `json:"id_user"`
	TypeTrk          int    `json:"type_trk"`
	StatusTransaksi  int    `json:"status_transaksi"`
	JumlahTrx        string `json:"jumlah_trx"`
	TrxType2         int    `json:"trx_type2"`
	StatusPembayaran int    `json:"status_pembayaran"`
	SellerId         int    `json:"seller_id"`
	FotoProduct      string `json:"foto_product"`
	NmPro            string `json:"nm_pro"`
	SatuanProduct    string `json:"satuan_product"`
	HargaSatuan1     string `json:"harga_satuan1"`
}

type TransactionHistory struct {
	// gorm.Model
	IdTransaksi      int    `json:"id_transaksi"`
	Date             string `json:"date"`
	Keterangan       string `json:"keterangan"`
	Balance          string `json:"balance"`
	NumberTrx        string `json:"number_trx"`
	IdProduct        int    `json:"id_product"`
	IdUser           int    `json:"id_user"`
	TypeTrk          int    `json:"type_trk"`
	StatusTransaksi  int    `json:"status_transaksi"`
	JumlahTrx        string `json:"jumlah_trx"`
	TrxType2         int    `json:"trx_type2"`
	StatusPembayaran int    `json:"status_pembayaran"`
	SellerId         int    `json:"seller_id"`
	TypeProduct      string `json:"type_product"`
	NmProduct        string `json:"nm_product"`
	Foto             string `json:"foto"`
}
