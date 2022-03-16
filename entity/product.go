package entity

type Product struct {
	IdProduct          int    `json:"id_product"`
	NamaProduct        string `json:"nama_product"`
	HargaSatuan        string `json:"harga_satuan"`
	IdKategori         string `json:"id_kategori"`
	DeskripsiProduct   string `json:"deskripsi_product"`
	BeratProduct       int    `json:"berat_product"`
	ManfaatProduct     string `json:"manfaat_product"`
	DateCreate         string `json:"date_create"`
	PenyimpananProduct string `json:"penyimpanan_product"`
	Img_1              string `json:"img_1"`
	Img_2              string `json:"img_2"`
	Img_3              string `json:"img_3"`
	Img_4              string `json:"img_4"`
	Img_5              string `json:"img_5"`
}
