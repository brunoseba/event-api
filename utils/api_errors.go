package utils

import "net/http"

type ApiError struct {
	Code    int    `json:code`
	Message string `json:message`
	//Cause   Cause  `json:cause`
}

var (
	ErrNotFound     = &ApiError{Code: http.StatusNotFound, Message: "Resourse no found"}
	ErrUnauthorized = &ApiError{Code: http.StatusUnauthorized, Message: "Access denegate"}
	ErrServerError  = &ApiError{Code: http.StatusBadGateway, Message: "Error with Server"}
)

func (e *ApiError) NewApiError(code int, message string) ApiError {
	return ApiError{
		Code:    code,
		Message: message,
	}
}
