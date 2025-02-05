package util

import "os"

// Getenv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment and non-empty, its value is returned.
// Otherwise, the fallback value is returned.
//
// Example:
//
//	os.Setenv("EXAMPLE_KEY", "example_value")
//	value := Getenv("EXAMPLE_KEY", "default_value")  // Returns "example_value"
//
//	value = Getenv("NON_EXISTING_KEY", "default_value")  // Returns "default_value"
func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func CheckString[T any](val T) bool {
	str, ok := any(val).(string)
	return ok && str != ""
}
