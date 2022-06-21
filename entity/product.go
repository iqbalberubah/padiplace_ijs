package entity

type Product struct {
	IdProduct        int    `json:"id_product"`
	NMProduct        string `json:"nm_product"`
	HargaProduct     string `json:"harga_product"`
	JmlProduct       string `json:"jml_product"`
	DeskripsiProduct string `json:"deskripsi_product"`
	BeratProduct     string `json:"berat_product"`
}
