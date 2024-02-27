package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rifatrakib/go-mechanic-shop/configuration"
	"github.com/rifatrakib/go-mechanic-shop/routes"
)

func main() {
	config := configuration.LoadConfig()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(config.CORSConfig))

	e.GET("/health", routes.HealthCheck)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Port)))
}
