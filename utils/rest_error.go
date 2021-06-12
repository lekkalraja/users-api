package utils

import (
	"fmt"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func UserNotFound(id int64) *RestErr {
	return &RestErr{
		Message: fmt.Sprintf("%d not found in the repo", id),
		Status:  http.StatusNoContent,
		Error:   "no_content",
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: fmt.Sprintf("something went wrong while processing request : %s", message),
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
