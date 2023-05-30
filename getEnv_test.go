package main

import (
	"testing"
)

type MockEnvLoader struct{}

func (m *MockEnvLoader) Load() error {
	return nil
}

func (m *MockEnvLoader) Process(prefix string, spec interface{}) error {
	config := spec.(*Config)
	config.Auth = "http://test.auth:8000"
	config.Back = "http://test.back:8001"
	config.Lobby = "http://test.lobby:8002"
	return nil
}

func Test_getEnv(t *testing.T) {
	mockLoader := &MockEnvLoader{}
	getEnv(mockLoader)

	if Env.Auth != "http://test.auth:8000" || Env.Back != "http://test.back:8001" || Env.Lobby != "http://test.lobby:8002" {
		t.Errorf("getEnv() error, expected environment variables to be set from mock")
	}
}
