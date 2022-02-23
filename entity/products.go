package entity

type Product struct {
	IdProduct          int    `json:"id_product"`
	NamaProduct        string `json:"nama_product"`
	IdKategori         string `json:"id_kategori"`
	HargaSatuan        string `json:"harga_satuan"`
	DeskripsiProduct   string `json:"deskripsi_product"`
	BeratProduct       string `json:"berat_product"`
	ManfaatProduct     string `json:"manfaat_product"`
	DateCreate         string `json:"date_create"`
	PenyimpananProduct string `json:"penyimpanan_product"`
	Img1               string `json:"img_1"`
	Img2               string `json:"img_2"`
	Img3               string `json:"img_3"`
	Img4               string `json:"img_4"`
	Img5               string `json:"img_5"`
}
