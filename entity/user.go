package entity

type User struct {
	// gorm.Model
	IdUser        int    `json:"id_user"`
	NameUser      string `json:"name_user"`
	EmailUser     string `json:"email_user"`
	ImgUser       string `json:"img_user"`
	Password1User string `json:"password1_user"`
	IdRole        int    `json:"id_role"`
	IsActive      int    `json:"is_active"`
	DateCreate    string `json:"date_create"`
	TlpUser       string `json:"tlp_user"`
	NpwpUser      string `json:"npwp_user"`
	JenisUsaha    string `json:"jenis_usaha"`
	AlamatUser    string `json:"alamat_user"`
	NamaPrusahaan string `json:"nama_prusahaan"`
	Username      string `json:"username"`
}
