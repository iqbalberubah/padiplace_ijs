package crud_po

import "padiplace_ijs/entity"

type UseCase interface {
	Get(id int) (*entity.PO, error)
	// GetAll() ([]*entity.PO, error)
	Create(in entity.PO) (*entity.PO, error)
	Update(in entity.PO) (*entity.PO, error)
	Delete(id int) error
}
