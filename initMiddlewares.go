package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initMiddlewares(echoApp *echo.Echo) {
	// echoApp.Use(middleware.Logger())
	echoApp.Use(middleware.Recover())
	echoApp.Static("/admin", "admin/dist")
	echoApp.GET("/hello", hello)
	echoApp.Use(gateway)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo!!!")
}
