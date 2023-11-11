package helper

import "net/http"

type Error interface {
	Message() string
	Status() int
	Type() string
}

type ErrorResponse struct {
	ErrorMessage string `json:"message"`
	ErrorStatus int `json:"status"`
	ErrorType string `json:"error"`
}

func (e *ErrorResponse) Message() string {
	return e.ErrorMessage
}

func (e *ErrorResponse) Status() int {
	return e.ErrorStatus
}

func (e *ErrorResponse) Type() string {
	return e.ErrorType
}

func NewError(message string, status int, errorType string) Error {
	return &ErrorResponse{
		ErrorMessage: message,
		ErrorStatus: status,
		ErrorType: errorType,
	}
}

func BadRequest(message string) Error {
	return NewError(message, http.StatusBadRequest, "Bad Request")
}

func Unauthorized(message string) Error {
	return NewError(message, http.StatusUnauthorized, "Unauthorized")
}

func NotFound(message string) Error {
	return NewError(message, http.StatusNotFound, "Not Found")
}

func UnprocessibleEntity(message string) Error {
	return NewError(message, http.StatusUnprocessableEntity, "Invalid Request")
}

func InternalServerError(message string) Error {
	return NewError(message, http.StatusInternalServerError, "Server Error")
}
