package main

type Product struct {
	// gorm.Model
	IdProduct          int    `json:"id_product"`
	NamaProduct        string `json:"nama_product"`
	HargaSatuan        string `json:"harga_satuan"`
	IdKategori         string `json:"id_kategori"`
	DeskripsiProduct   string `json:"deskripsi_product"`
	BeratProduct       int    `json:"berat_product"`
	ManfaatProduct     string `json:"manfaat_product"`
	DateCreate         bool   `json:"date_create"`
	PenyimpananProduct bool   `json:"penyimpanan_product"`
	Img1               bool   `json:"img_1"`
	Img2               bool   `json:"img_2"`
	Img3               bool   `json:"img_3"`
	Img4               bool   `json:"img_4"`
	Img5               bool   `json:"img_5"`
}
