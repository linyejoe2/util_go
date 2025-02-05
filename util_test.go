package util_test

import (
	"os"
	"testing"

	util "github.com/linyejoe2/util_go"
)

func TestGetenv(t *testing.T) {
	// Test case 1: Environment variable is set
	os.Setenv("EXISTING_KEY", "value")
	result := util.Getenv("EXISTING_KEY", "fallback")
	if result != "value" {
		t.Errorf("expected 'value', got '%s'", result)
	}

	// Test case 2: Environment variable is not set, fallback is returned
	os.Unsetenv("NON_EXISTING_KEY")
	result = util.Getenv("NON_EXISTING_KEY", "fallback")
	if result != "fallback" {
		t.Errorf("expected 'fallback', got '%s'", result)
	}

	// Test case 3: Empty environment variable is set, fallback is returned
	os.Setenv("EMPTY_KEY", "")
	result = util.Getenv("EMPTY_KEY", "fallback")
	if result != "fallback" {
		t.Errorf("expected 'fallback', got '%s'", result)
	}
}
