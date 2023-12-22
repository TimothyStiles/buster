package test

import "testing"

func TestTest(t *testing.T) {
	testCommand := []string{"go", "test", "./test/mock/mock_test.go"}
	results, err := Test(testCommand)

	if err != nil {
		t.Errorf("Test() = %v, want nil", err)
	}

	if results == "" {
		t.Errorf("Test() = %v, want nil", results)
	}

	testCommand = []string{"./test/mock/mock_test.go"}
	results, err = Test(testCommand)

	if err != nil {
		t.Errorf("Test() = %v, want nil", err)
	}

	if results == "" {
		t.Errorf("Test() = %v, want nil", results)
	}

	testCommand = []string{"-v", "./test/mock/mock_test.go"}
	results, err = Test(testCommand)

	if err != nil {
		t.Errorf("Test() = %v, want nil", err)
	}

	if results == "" {
		t.Errorf("Test() = %v, want nil", results)
	}
}
