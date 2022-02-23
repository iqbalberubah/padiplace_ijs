package crud_product

import "padiplace_ijs/entity"

type UseCase interface {
	Get(id int) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Create(in entity.Product) (*entity.Product, error)
	Update(in entity.Product) (*entity.Product, error)
	Delete(id int) error
}
