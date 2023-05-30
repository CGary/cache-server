package environment

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Auth  string `default:"http://100.12.212.45:8000"`
	Back  string `default:"http://100.12.212.45:8001"`
	Lobby string `default:"http://100.12.212.45:8002"`
}

func (c *Config) getEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	err = envconfig.Process("CS", c)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
