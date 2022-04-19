package tests

import (
	"testing"

	"proxy-collect/config"
)

func TestConfig(t *testing.T) {
	config.LoadConfig()
	t.Log(config.Get())
}
