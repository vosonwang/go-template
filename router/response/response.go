package response

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Style struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Style{
		Success: true,
		Data:    data,
	})
}

func Fail(c echo.Context, message interface{}) error {
	return c.JSON(http.StatusOK, Style{
		Message: fmt.Sprintf("%v", message),
	})
}

func NotFound(c echo.Context) error {
	return c.JSON(http.StatusNotFound, Style{
		Message: http.StatusText(http.StatusNotFound),
	})
}

func BadRequest(c echo.Context, message interface{}) error {
	return c.JSON(http.StatusBadRequest, Style{
		Message: fmt.Sprintf("%v", message),
	})
}

func Forbidden(c echo.Context) error {
	return c.JSON(http.StatusForbidden, Style{
		Message: http.StatusText(http.StatusForbidden),
	})
}

func ServiceUnavailable(c echo.Context, message interface{}) error {
	return c.JSON(http.StatusServiceUnavailable, Style{
		Message: fmt.Sprintf("%v", message),
	})
}
