package product_repository

import (
	"errors"
	"padiplace_ijs/entity"

	irepository "padiplace_ijs/app/repository"

	"gorm.io/gorm"
)

type repository struct {
	db        *gorm.DB
	tableName string
}

func NewProduct(db *gorm.DB) irepository.ProductRepository {
	return &repository{db, "product"}
}

func (r *repository) Get(id int) (*entity.Product, error) {
	resp := ProductModel{}
	err := r.db.Table(r.tableName).Where("id_product = ?", id).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, irepository.ErrProductNotFound
		}
		return nil, err
	}
	return resp.ToProductEntity(), nil
}

func (r *repository) GetAll() ([]*entity.Product, error) {
	datas := []ProductModel{}
	err := r.db.Table(r.tableName).Find(&datas).Error
	if err != nil {
		return nil, err
	}
	resp := []*entity.Product{}

	for _, data := range datas {
		resp = append(resp, data.ToProductEntity())
	}
	return resp, nil

}
func (r *repository) Create(in entity.Product) (*entity.Product, error) {
	productModel := ProductModel{}.FromProductEntity(in)

	// timeNow := time.Now()
	// productModel.CreatedAt = &timeNow
	// productModel.UpdatedAt = &timeNow

	err := r.db.Table(r.tableName).Create(&productModel).Error
	if err != nil {
		return nil, err
	}
	return productModel.ToProductEntity(), nil

}
func (r *repository) Update(in entity.Product) (*entity.Product, error) {
	productModel := ProductModel{}.FromProductEntity(in)
	_, err := r.Get(in.IdProduct)
	if errors.Is(err, irepository.ErrProductNotFound) {
		return nil, nil
	}

	// timeNow := time.Now()
	// productModel.CreatedAt = nil
	// productModel.UpdatedAt = &timeNow
	err = r.db.Table(r.tableName).Where("id_product = ?", in.IdProduct).Updates(&productModel).Error
	if err != nil {
		return nil, err
	}

	return productModel.ToProductEntity(), nil

}
func (r *repository) Delete(id int) error {
	return r.db.Table(r.tableName).Delete(&ProductModel{}, "id_product = ?", id).Error
}
