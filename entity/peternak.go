package entity

type Peternak struct {
	// gorm.Model
	IdPeternak     int    `json:"id_peternak"`
	NamaPeternak   string `json:"nama_peternak"`
	AlamatPeternak string `json:"alamat_peternak"`
	ImgPeternak    string `json:"ImgPeternak"`
	ImgKtp         string `json:"img_ktp"`
	StatusPeternak string `json:"status_peternak"`
	TlpPeternak    string `json:"tlp_peternak"`
	IsActive       int    `json:"is_active"`
	Password1User  string `json:"password1_user"`
}
