package response

import (
	"sistem-manajemen-armada/api/dto"
)

func GenerateSuccessResponse(message string, data interface{}) dto.SuccessResponse {
	return dto.SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
