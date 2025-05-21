package dto

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FailedResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}