package tester_test

import (
	"testing"

	"github.com/TimothyStiles/buster/tester"
)

// func TestTester(t *testing.T) {

// 	// testConfig := tester.Config{
// 	// 	rootPath: "../",
// 	// }
// 	results, err := tester.Tester()
// 	if err != nil {
// 		t.Errorf("Expected no errors, got %s", err)
// 	}

// 	if results != "" {
// 		t.Errorf("Expected no errors, got %s", results)
// 	}
// }

func TestRunTests(t *testing.T) {
	if err := tester.RunTests("./mock"); err != nil {
		t.Errorf("Expected no errors, got %s", err)
	}
}

func TestRunTestContainer(t *testing.T) {
	if err := tester.RunTestContainer("./tester/mock"); err != nil {
		t.Errorf("Expected no errors, got %s", err)
	}
}
