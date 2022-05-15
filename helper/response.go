package helper

import (
	"strings"
)

// Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// EmptyObj object is used when the data is empty
type EmptyObj struct{}


// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}

	return res
}

// BuildErrorResponse method is to inject error value to dynamic error response
func BuildErrorResponse(mmessage string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")

	res := Response{
		Status:  false,
		Message: mmessage,
		Error:   splittedError,
		Data:    nil,
	}

	return res
}
