package util

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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

// CheckString checks if the input is a string and not an empty string.
// It returns true if the input is a non-empty string, otherwise false.
//
// Example:
//
//	result := CheckString("hello")   // returns true
//	result := CheckString("")        // returns false
//	result := CheckString(123)       // returns false
func CheckString[T any](val T) bool {
	str, ok := any(val).(string)
	return ok && str != ""
}

// ToInt converts a value of any numeric type to an integer.
// It handles signed and unsigned integers, as well as floating-point numbers.
// If the conversion results in overflow, it returns an error.
func ToInt(val interface{}) int {
	switch v := any(val).(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		if v >= uint64(^uint(0)) { // check for overflow
			panic(fmt.Sprintf("value %v overflows int", v))
		}
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	default:
		panic(fmt.Sprintf("cannot convert %v to int", v))
	}
}

func ResponseBadRequest(c *gin.Context, message string, body interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": message, "body": body})
}

func ResponseCustom(c *gin.Context, code int, error bool, message string, body interface{}) {
	c.JSON(code, gin.H{"error": true, "message": message, "body": body})
}

func ResponseOK(c *gin.Context, message string, body interface{}) {
	c.JSON(http.StatusOK, gin.H{"error": false, "message": message, "body": body})
}

// DotProduct calculates the dot product of two vectors.
func DotProduct[T float32 | float64](a, b []T) T {
	if len(a) != len(b) {
		panic("vectors must have the same length")
	}

	var sum T
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}
