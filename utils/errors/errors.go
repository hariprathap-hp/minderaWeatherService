package errors

import "net/http"

type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   message,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Message: "Bad Request Error",
		Error:   message,
	}
}
