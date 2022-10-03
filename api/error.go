package api

import "net/http"

var (
	ErrInternalError = NewError("Internal Server Error", http.StatusInternalServerError)
	ErrUnauthorized  = NewError("Unauthorized", http.StatusUnauthorized)
)

func NewError(message string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
