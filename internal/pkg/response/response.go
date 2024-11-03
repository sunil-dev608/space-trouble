package response

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// SuccessResponse returns a success response
func SuccessResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse returns an error response
func ErrorResponse(c echo.Context, status int, message string, err interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Message: message,
		Error:   err,
	})
}
