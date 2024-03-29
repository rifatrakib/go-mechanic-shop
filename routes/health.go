package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{Status: "healthy"})
}
