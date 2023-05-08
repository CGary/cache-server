package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initApiServer() {
	echoApp := echo.New()
	echoApp.Use(middleware.Logger())
	echoApp.Use(middleware.Recover())
	mainRoutes((echoApp))
	echoApp.Logger.Fatal(echoApp.Start(":8081"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hola mundo!!!")
}

func mainRoutes(echoApp *echo.Echo) {
	echoApp.GET("/", hello)
}
