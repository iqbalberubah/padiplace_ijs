package entity

type Balance struct {
	// gorm.Model
	IdBalance int    `json:"id_balance"`
	Balance   string `json:"balance"`
	IdKios    string `json:"id_kios"`
	Trx       string `json:"trx"`
}
