package repository

import (
	"errors"
	"padiplace_ijs/entity"
)

var ErrProductNotFound = errors.New("product not found")

type VideoRepository interface {
	Get(id int) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Create(in entity.Product) (*entity.Product, error)
	Update(in entity.Product) (*entity.Product, error)
	Delete(id int) error
}
