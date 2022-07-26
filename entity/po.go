package entity

type PO struct {
	// gorm.Model
	IdPo         int    `json:"id_po"`
	NomerPo      string `json:"nomer_po"`
	Jumlah       string `json:"jumlah"`
	DateCreate   string `json:"date_create"`
	Status       int    `json:"status"`
	Total        string `json:"total"`
	HargaSatuan  string `json:"harga_satuan"`
	IdProduct    int    `json:"id_product"`
	IdKios       int    `json:"id_kios"`
	IdPengiriman int    `json:"id_pengiriman"`
	AlamatPO     int    `json:"alamat_po"`
	Lang         int    `json:"lang"`
	Lat          int    `json:"lat"`
}
