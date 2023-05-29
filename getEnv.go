package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Auth  string `default:"http://100.12.212.45:8000"`
	Back  string `default:"http://100.12.212.45:8001"`
	Lobby string `default:"http://100.12.212.45:8002"`
}

var Env Config

func getEnv() {
	godotenv.Load()
	envconfig.Process("CS", &Env)
}
