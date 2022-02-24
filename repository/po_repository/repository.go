package po_repository

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

func NewPO(db *gorm.DB) irepository.PORepository {
	return &repository{db, "po"}
}

func (r *repository) Get(id int) (*entity.PO, error) {
	resp := POModel{}
	err := r.db.Table(r.tableName).Where("id_po = ?", id).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, irepository.ErrPONotFound
		}
		return nil, err
	}
	return resp.ToPOEntity(), nil
}

// func (r *repository) GetAll() ([]*entity.PO, error) {
// 	datas := []POModel{}
// 	err := r.db.Table(r.tableName).Find(&datas).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp := []*entity.PO{}

// 	for _, data := range datas {
// 		resp = append(resp, data.ToPOEntity())
// 	}
// 	return resp, nil

// }

func (r *repository) Create(in entity.PO) (*entity.PO, error) {
	poModel := POModel{}.FromPOEntity(in)

	// timeNow := time.Now()
	// productModel.CreatedAt = &timeNow
	// productModel.UpdatedAt = &timeNow

	err := r.db.Table(r.tableName).Create(&poModel).Error
	if err != nil {
		return nil, err
	}
	return poModel.ToPOEntity(), nil

}
func (r *repository) Update(in entity.PO) (*entity.PO, error) {
	poModel := POModel{}.FromPOEntity(in)
	_, err := r.Get(in.IdPO)
	if errors.Is(err, irepository.ErrPONotFound) {
		return nil, nil
	}

	// timeNow := time.Now()
	// productModel.CreatedAt = nil
	// productModel.UpdatedAt = &timeNow
	err = r.db.Table(r.tableName).Where("id_po = ?", in.IdPO).Updates(&poModel).Error
	if err != nil {
		return nil, err
	}

	return poModel.ToPOEntity(), nil

}
func (r *repository) Delete(id int) error {
	return r.db.Table(r.tableName).Delete(&POModel{}, "id_po = ?", id).Error
}
