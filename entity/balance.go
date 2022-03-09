package entity

type Balance struct {
	// gorm.Model
	IdBalance  int    `json:"id_balance"`
	Balance    string `json:"balance"`
	IdPeternak string `json:"id_peternak"`
	Trx        string `json:"trx"`
}
