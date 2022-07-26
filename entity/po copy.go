package entity

type Stock struct {
	// gorm.Model
	IdStock     int    `json:"id_stock"`
	IdKios      int    `json:"id_kios"`
	JumlahStock string `json:"jumlah_stock"`
	IdProduct   int    `json:"id_product"`
}
