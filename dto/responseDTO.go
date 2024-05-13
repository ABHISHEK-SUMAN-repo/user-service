package dto

type ResponseDTO struct {
	Status         bool        `json:"status"`
	Message        string      `json:"message"`
	Id             string      `json:"id"`
	Data           interface{} `json:"data"`
	ErrorCode      int         `json:"error"`
	HttpStatusCode int         `json:"http_status_code"`
}

func Success() ResponseDTO {
	return ResponseDTO{
		Status:         true,
		Message:        "Success",
		HttpStatusCode: 200,
	}
}

func Faliure() ResponseDTO {
	return ResponseDTO{
		Status:         false,
		Message:        "Faliure",
		HttpStatusCode: 500,
	}
}

func SuccessWithData(data interface{}) ResponseDTO {
	return ResponseDTO{
		Status:         true,
		Message:        "Success",
		Data:           data,
		HttpStatusCode: 200,
	}
}

func FailedWithData(data interface{}) ResponseDTO {
	return ResponseDTO{
		Status:         false,
		Message:        "Faliure",
		Data:           data,
		HttpStatusCode: 500,
	}
}
