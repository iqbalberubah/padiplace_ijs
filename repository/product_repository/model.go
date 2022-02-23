package product_repository

import (
	"padiplace_ijs/entity"
)

type ProductModel struct {
	// Id        int        `gorm:"primary_key;column:id"`
	IdProduct        int    `gorm:"primary_key;column:id_product"`
	NamaProduct      string `gorm:"column:nama_product"`
	DeskripsiProduct string `gorm:"column:deskripsi_product"`
	// HargaSatuan *time.Time `gorm:"column:created_at"`
	// UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (ProductModel) FromProductEntity(v entity.Product) *ProductModel {
	return &ProductModel{
		IdProduct:        v.IdProduct,
		NamaProduct:      v.NamaProduct,
		DeskripsiProduct: v.DeskripsiProduct,
	}
}

func (m *ProductModel) ToProductEntity() *entity.Product {
	return &entity.Product{
		IdProduct:        m.IdProduct,
		NamaProduct:      m.NamaProduct,
		DeskripsiProduct: m.DeskripsiProduct,
	}
}
