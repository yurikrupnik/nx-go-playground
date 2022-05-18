package my_lib

import (
	"testing"
)

func TestMyLib(t *testing.T) {
	result := MyLib("works")
	if result != "MyLib works" {
		t.Error("Expected MyLib to append 'works'")
	}
}
