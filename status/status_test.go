package status

import (
	"testing"
)

func TestCheck(t *testing.T) {
	if err := Check(); err != nil {
		t.Errorf("Check() = %v, want nil", err)
	}
}
