package environment

import (
	"testing"
)

type MockConfig struct {
	Auth  string `default:"http://100.12.212.45:8000"`
	Back  string `default:"http://100.12.212.45:8001"`
	Lobby string `default:"http://100.12.212.45:8002"`
}

const auth = "http://test.auth:8000"
const back = "http://test.back:8001"
const lobby = "http://test.lobby:8002"

func (c *MockConfig) getEnv() error {
	c.Auth = auth
	c.Back = back
	c.Lobby = lobby
	return nil
}

func Test_getEnv(t *testing.T) {
	mockLoader := &MockConfig{}
	err := mockLoader.getEnv()
	if mockLoader.Auth != auth || mockLoader.Back != back || mockLoader.Lobby != lobby {
		t.Errorf("getEnv() error, expected environment variables to be set from mock")
	}
	if err != nil {
		t.Error("Expected error, got nil")
	}
}
