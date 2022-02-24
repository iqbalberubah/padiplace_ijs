package repository

import (
	"errors"
	"padiplace_ijs/entity"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	Get(id int) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Create(in entity.Product) (*entity.Product, error)
	Update(in entity.Product) (*entity.Product, error)
	Delete(id int) error
}

var ErrPONotFound = errors.New("po not found")

type PORepository interface {
	Get(id int) (*entity.PO, error)
	Create(in entity.PO) (*entity.PO, error)
	Update(in entity.PO) (*entity.PO, error)
	Delete(id int) error
	// GetAll() ([]*entity.PO, error)
}
