package lint_test

import (
	"testing"

	"github.com/TimothyStiles/buster/config"
	lint "github.com/TimothyStiles/buster/linter"
)

func TestRunLint(t *testing.T) {
	cfg := config.NewConfig()
	cfg.RootPath = "../"

	_, err := lint.RunLint(cfg, "./tester/mock")
	if err != nil {
		t.Errorf("Expected no errors, got %s", err)
	}
}
