package util_test

import (
	"fmt"
	"os"
	"testing"

	util "github.com/linyejoe2/util_go"
	"github.com/stretchr/testify/assert"
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

func TestCheckString(t *testing.T) {
	// Test case 1: Valid non-empty string
	result := util.CheckString("hello")
	if !result {
		t.Errorf("expected true, got false")
	}

	// Test case 2: Empty string
	result = util.CheckString("")
	if result {
		t.Errorf("expected false, got true")
	}

	// Test case 3: Non-string input (int)
	result = util.CheckString(123)
	if result {
		t.Errorf("expected false, got true")
	}

	// Test case 4: Non-string input (struct)
	type ExampleStruct struct{}
	result = util.CheckString(ExampleStruct{})
	if result {
		t.Errorf("expected false, got true")
	}

	// Test case 5: String with spaces
	result = util.CheckString("   ")
	if !result {
		t.Errorf("expected true, got false")
	}

	// Test case 6: Non-string slice
	result = util.CheckString([]int{1, 2, 3})
	if result {
		t.Errorf("expected false, got true")
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    int
		wantErr bool
	}{
		{"int", int(10), 10, false},
		{"int8", int8(20), 20, false},
		{"int16", int16(30), 30, false},
		{"int32", int32(40), 40, false},
		{"int64", int64(50), 50, false},
		{"uint", uint(60), 60, false},
		{"uint8", uint8(70), 70, false},
		{"uint16", uint16(80), 80, false},
		{"uint32", uint32(90), 90, false},
		{"uint64 overflow", uint64(1<<64 - 1), 0, true}, // Overflow case
		{"float32", float32(1.5), 1, false},
		{"float64", float64(2.7), 2, false},
		{"unsupported type", "invalid", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				defer func() {
					r := recover()
					if r == nil {
						t.Errorf("expected panic on %v, but got none", tt.name)
					} else {
						fmt.Println("Recover: ", r)
					}
				}()
				_ = util.ToInt(tt.input)
			} else {
				got := util.ToInt(tt.input)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
