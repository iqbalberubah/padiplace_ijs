package entity

type TransactionDetail struct {
	// gorm.Model
	IdInv           int    `json:"id_inv"`
	NomerTrx        string `json:"nomer_trx"`
	Tanggal         string `json:"tanggal"`
	IdSeller        int    `json:"id_seller"`
	IdBuyer         int    `json:"id_buyer"`
	IdPengiriman    int    `json:"id_pengiriman"`
	TotalPemesanan  string `json:"total_pemesanan"`
	JumlahPro       string `json:"jumlah_pro"`
	StatusTransaksi int    `json:"status_transaksi"`
	TrxType2        int    `json:"trx_type2"`
	Nomerresi       string `json:"nomerresi"`
	Pengiriman      string `json:"pengiriman"`
}
