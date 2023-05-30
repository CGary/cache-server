package main

import "github.com/labstack/echo"

func main() {
	envLoader := &EnvConfigLoader{}
	getEnv(envLoader)
	echoApp := echo.New()
	initMiddlewares(echoApp)
	echoApp.Logger.Fatal(echoApp.Start(":8081"))
}
