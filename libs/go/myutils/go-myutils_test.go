package go_myutils

import (
	"testing"
)

func TestMyutils(t *testing.T) {
	result := Myutils("works")
	if result != "Myutils works" {
		t.Error("Expected Myutils to append 'works'")
	}
}
