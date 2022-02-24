package crud_product

import (
	"errors"
	"fmt"
	"padiplace_ijs/app/repository"
	"padiplace_ijs/entity"
)

type usecase struct {
	product_repo repository.ProductRepository
}

func NewUseCase(product_repo repository.ProductRepository) UseCase {
	return &usecase{product_repo: product_repo}
}

func (uc *usecase) Get(id int) (*entity.Product, error) {
	data, err := uc.product_repo.Get(id)
	if err != nil {
		if errors.Is(err, repository.ErrProductNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) GetAll() ([]*entity.Product, error) {
	data, err := uc.product_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}
func (uc *usecase) Create(in entity.Product) (*entity.Product, error) {
	data, err := uc.product_repo.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) Update(in entity.Product) (*entity.Product, error) {
	data, err := uc.product_repo.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil

}

func (uc *usecase) Delete(id int) error {
	err := uc.product_repo.Delete(id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return err
}
