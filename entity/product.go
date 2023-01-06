package entity

type Product struct {
	IdProduct       int    `json:"id_product"`
	NmProduct       string `json:"nm_product"`
	HargaProduct    string `json:"harga_product"`
	Stock           string `json:"stock"`
	TypeProduct     string `json:"type_product"`
	SatuanProduct   string `json:"satuan_product"`
	DesProduct      string `json:"des_product"`
	Foto            string `json:"foto"`
	Foto2           string `json:"foto2"`
	Foto3           string `json:"foto3"`
	Foto4           string `json:"foto4"`
	Foto5           string `json:"foto5"`
	Foto6           string `json:"foto6"`
	Status          int    `json:"status"`
	IdUser          int    `json:"id_user"`
	TglDate         string `json:"tgl_date"`
	KomProduct      string `json:"kom_product"`
	TipsPengProduct string `json:"tips_peng_product"`
	NoSku           string `json:"no_sku"`
	PublishS        int    `json:"publish_s"`
}

type ProductList struct {
	IdProduct       int    `json:"id_product"`
	NmProduct       string `json:"nm_product"`
	HargaProduct    string `json:"harga_product"`
	Stock           string `json:"stock"`
	TypeProduct     string `json:"type_product"`
	SatuanProduct   string `json:"satuan_product"`
	DesProduct      string `json:"des_product"`
	Foto            string `json:"foto"`
	Foto2           string `json:"foto2"`
	Foto3           string `json:"foto3"`
	Foto4           string `json:"foto4"`
	Foto5           string `json:"foto5"`
	Foto6           string `json:"foto6"`
	Status          int    `json:"status"`
	IdUser          int    `json:"id_user"`
	TglDate         string `json:"tgl_date"`
	KomProduct      string `json:"kom_product"`
	TipsPengProduct string `json:"tips_peng_product"`
	NoSku           string `json:"no_sku"`
	PublishS        int    `json:"publish_s"`
	NamaPrusahaan   string `json:"nama_prusahaan"`
}
