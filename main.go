package main

import "github.com/labstack/echo"

func main() {
	getEnv()
	echoApp := echo.New()
	initMiddlewares(echoApp)
	echoApp.Logger.Fatal(echoApp.Start(":8081"))
}
