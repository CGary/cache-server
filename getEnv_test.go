package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("CS_AUTH", "http://localhost:8000")
	defer os.Unsetenv("CS_AUTH")
	getEnv()
	if Env.Auth != "http://localhost:8000" {
		t.Errorf("getEnv() failed, expected %s but got %s", "http://localhost:8000", Env.Auth)
	}
}

func TestGetEnvWithDotEnv(t *testing.T) {
	err := godotenv.Load("example.env")
	if err != nil {
		t.Errorf("Error loading example.env file: %v", err)
	}
	getEnv()
	if Env.Auth != "http://localhost:8000" {
		t.Errorf("getEnv() failed, expected %s but got %s", "http://localhost:8000", Env.Auth)
	}
}
