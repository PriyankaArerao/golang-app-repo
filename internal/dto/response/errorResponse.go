package response

import (
	"fmt"
)

type ErrorResponse struct {
	ErrorCode  string `json:"code"`
	Message    string `json:"msg"`
	StatusCode int    `json:"-"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%d: (%s) %s", e.StatusCode, e.ErrorCode, e.Message)
}
