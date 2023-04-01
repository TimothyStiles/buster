package tester_test

import (
	"testing"

	"github.com/TimothyStiles/buster/config"
	"github.com/TimothyStiles/buster/tester"
)

func TestRunTestContainer(t *testing.T) {
	cfg := config.NewConfig()
	cfg.RootPath = "../"

	_, err := tester.RunTestContainer(cfg, "./tester/mock")
	if err != nil {
		t.Errorf("Expected no errors, got %s", err)
	}
}
