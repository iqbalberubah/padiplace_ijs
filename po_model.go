package main

type PO struct {
	// gorm.Model
	IdPo        int    `json:"id_po"`
	NomerPo     string `json:"nomer_po"`
	Jumlah      string `json:"jumlah"`
	DateCreate  string `json:"date_create"`
	IdClient    int    `json:"id_client"`
	Status      int    `json:"status"`
	Total       string `json:"total"`
	HargaSatuan string `json:"harga_satuan"`
	T           int    `json:"t"`
	Product     string `json:"product"`
	Box         string `json:"box"`
	Cp          string `json:"cp"`
	Stk         string `json:"stk"`
	IdPeternak  string `json:"id_peternak"`
	Numberpo    int    `json:"numberpo"`
	Popro       int    `json:"popro"`
	Popro1      int    `json:"popro1"`
}

type SuccesPO struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Data            PO     `json:"data"`
}

type SuccesPOData struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Data            []PO   `json:"data"`
}

type ErrorPO struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
