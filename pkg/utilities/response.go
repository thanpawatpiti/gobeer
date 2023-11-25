package utilities

type ResponseStruct struct {
	Status bool        `json:"status"`
	Code   interface{} `json:"code"`
	Data   interface{} `json:"data"`
}

func Response(status bool, code any, data interface{}) ResponseStruct {
	return ResponseStruct{
		Status: status,
		Code:   code,
		Data:   data,
	}
}
