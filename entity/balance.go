package entity

type Balance struct {
	// gorm.Model
	IdBalance  int    `json:"id_balance"`
	Balance    string `json:"balance"`
	IdPeternak string `json:"id_peternak"`
	Trx        string `json:"trx"`
}

type SuccesBalance struct {
	ResponseCode    int     `json:"response_code"`
	ResponseMessage string  `json:"response_message"`
	Data            Balance `json:"data"`
}

type ErrorBalance struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
