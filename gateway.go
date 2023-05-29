package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func gateway(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		// fmt.Printf("%+v\n", c.Request())
		var authorizedFlag = c.Path() == "/admin" || c.Path() == "/admin/*" || c.Path() == "/hello"
		apiKey := c.Request().Header.Get("apikey")
		if !authorizedFlag && apiKey == "" {
			return c.String(http.StatusUnauthorized, "apikey is missing")
		}
		return next(c)
	}
}
