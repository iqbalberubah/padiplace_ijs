package entity

type Balance struct {
	// gorm.Model
	IdUser  int    `json:"id_balance"`
	Balance string `json:"balance"`
}
