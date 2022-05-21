package go_myutils

import (
	"os"
)

func Myutils(name string) string {
	result := "Myutils " + name
	return result
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
