package response

import (
	"sistem-manajemen-armada/api/dto"
)

func GenerateFailedResponse(message string, err error) dto.FailedResponse {
	var errorDetail interface{}
	if err != nil {
		errorDetail = err.Error()
	} else {
		errorDetail = nil
	}

	return dto.FailedResponse{
		Status:  "failed",
		Message: message,
		Error:   errorDetail,
	}
}
