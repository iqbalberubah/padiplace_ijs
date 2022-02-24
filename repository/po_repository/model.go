package po_repository

import (
	"padiplace_ijs/entity"
)

type POModel struct {
	IdPO        int    `gorm:"primary_key;column:id_po"`
	NomerPO     string `gorm:"column:nomer_po"`
	Jumlah      string `gorm:"column:jumlah"`
	DateCreate  string `gorm:"column:date_create"`
	IdClient    int    `gorm:"column:id_client"`
	Status      int    `gorm:"column:status"`
	Total       string `gorm:"column:total"`
	HargaSatuan string `gorm:"column:harga_satuan"`
	T           int    `gorm:"column:t"`
	Product     string `gorm:"column:product"`
	Box         string `gorm:"column:box"`
	CP          string `gorm:"column:cp"`
	STK         int    `gorm:"column:stk"`
	Id_Peternak string `gorm:"column:id_peternak"`
	Numberpo    int    `gorm:"column:numberpo"`
	Popro       int    `gorm:"column:popro"`
	Popro1      int    `gorm:"column:popro1"`
	// HargaSatuan *time.Time `gorm:"column:created_at"`
	// UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

func (POModel) FromPOEntity(v entity.PO) *POModel {
	return &POModel{
		IdPO:        v.IdPO,
		NomerPO:     v.NomerPO,
		Jumlah:      v.Jumlah,
		DateCreate:  v.DateCreate,
		IdClient:    v.IdClient,
		Status:      v.Status,
		Total:       v.Total,
		HargaSatuan: v.HargaSatuan,
		T:           v.T,
		Product:     v.Product,
		Box:         v.Box,
		CP:          v.CP,
		STK:         v.STK,
		Id_Peternak: v.Id_Peternak,
		Numberpo:    v.Numberpo,
		Popro:       v.Popro,
		Popro1:      v.Popro1,
	}
}

func (m *POModel) ToPOEntity() *entity.PO {
	return &entity.PO{
		IdPO:        m.IdPO,
		NomerPO:     m.NomerPO,
		Jumlah:      m.Jumlah,
		DateCreate:  m.DateCreate,
		IdClient:    m.IdClient,
		Status:      m.Status,
		Total:       m.Total,
		HargaSatuan: m.HargaSatuan,
		T:           m.T,
		Product:     m.Product,
		Box:         m.Box,
		CP:          m.CP,
		STK:         m.STK,
		Id_Peternak: m.Id_Peternak,
		Numberpo:    m.Numberpo,
		Popro:       m.Popro,
		Popro1:      m.Popro1,
	}
}
