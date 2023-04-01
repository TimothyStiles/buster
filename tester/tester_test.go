package tester_test

import (
	"testing"

	"github.com/TimothyStiles/buster/config"
	"github.com/TimothyStiles/buster/tester"
)

func TestRunTests(t *testing.T) {
	cfg := config.NewConfig()
	cfg.RootPath = "../"

	_, err := tester.RunTests(cfg, "./tester/mock")
	if err != nil {
		t.Errorf("Expected no errors, got %s", err)
	}
}
