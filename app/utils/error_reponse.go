package utils

type ErrorResponse struct {
	Message string `json:"description"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
	}
}
