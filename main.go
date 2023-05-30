package main

import (
	"github.com/labstack/echo"
)

func main() {
	// envLoader := &Config{}
	// errEnv := envLoader.GetEnv()
	// if errEnv != nil {
	// 	return
	// }
	echoApp := echo.New()
	initMiddlewares(echoApp)
	echoApp.Logger.Fatal(echoApp.Start(":8081"))
}
