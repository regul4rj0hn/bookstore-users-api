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

func Conflict(data string) *Response {
	return &Response{
		Message: "conflict",
		Status:  http.StatusConflict,
		Data:    data,
	}
}

func NotFound(data string) *Response {
	return &Response{
		Message: "not_found",
		Status:  http.StatusNotFound,
		Data:    data,
	}
}
