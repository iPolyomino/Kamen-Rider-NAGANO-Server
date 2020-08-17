package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ErrBadRequest          = "Bad Request"
	ErrInternalServerError = "Internal Server Error"
)

type apiErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func apiResponseOK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func apiResponseError(c echo.Context, status int, message string, err error) error {
	log.Println(err)
	return c.JSON(status, apiErrorResponse{Code: fmt.Sprintf("%d", status), Message: message})
}

func apiResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, apiErrorResponse{Code: fmt.Sprintf("%d", status), Message: message})
}
