package entity

type SuccesResponse struct {
	ResponseCode    int         `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	Data            interface{} `json:"data"`
}

type SuccesResponseDatas struct {
	ResponseCode    int           `json:"response_code"`
	ResponseMessage string        `json:"response_message"`
	Data            []interface{} `json:"data"`
}

type ErrorResponse struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
