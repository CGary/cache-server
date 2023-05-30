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

type EnvLoader interface {
	Load() error
	Process(prefix string, spec interface{}) error
}

type EnvConfigLoader struct{}

func (e *EnvConfigLoader) Load() error {
	return godotenv.Load()
}

func (e *EnvConfigLoader) Process(prefix string, spec interface{}) error {
	return envconfig.Process(prefix, spec)
}

func getEnv(loader EnvLoader) {
	loader.Load()
	loader.Process("CS", &Env)
}
