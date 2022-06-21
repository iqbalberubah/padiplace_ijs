package entity

type Kios struct {
	// gorm.Model
	IdKios     int    `json:"id_kios"`
	NMKios     string `json:"nm_kios"`
	AlamatKios string `json:"alamat_kios"`
	PhotoKios  string `json:"photo_kios"`
	Username   string `json:"username"`
	TlpKios    string `json:"tlp_kios"`
	Password   string `json:"password"`
}
