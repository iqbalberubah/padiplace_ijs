package crud_po

import (
	"errors"
	"fmt"
	"padiplace_ijs/app/repository"
	"padiplace_ijs/entity"
)

type usecase struct {
	po_repo repository.PORepository
}

func NewUseCase(po_repo repository.PORepository) UseCase {
	return &usecase{po_repo: po_repo}
}

func (uc *usecase) Get(id int) (*entity.PO, error) {
	data, err := uc.po_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

// func (uc *usecase) GetAll() ([]*entity.PO, error) {
// 	data, err := uc.po_repo.GetAll()
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
// 	}
// 	return data, nil
// }

func (uc *usecase) Create(in entity.PO) (*entity.PO, error) {
	data, err := uc.po_repo.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) Update(in entity.PO) (*entity.PO, error) {
	data, err := uc.po_repo.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil

}

func (uc *usecase) Delete(id int) error {
	err := uc.po_repo.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return err
}
