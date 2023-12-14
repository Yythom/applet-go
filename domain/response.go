package domain

type Response struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
