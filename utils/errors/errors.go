package errors

import "net/http"

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    string `json:"data"`
}

func BadRequest(data string) *Response {
	return &Response{
		Message: "bad_request",
		Status:  http.StatusBadRequest,
		Data:    data,
	}
}
