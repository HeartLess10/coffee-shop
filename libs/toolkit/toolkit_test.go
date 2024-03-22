package toolkit

import (
	"testing"
)

func TestToolkit(t *testing.T) {
	result := Toolkit("works")
	if result != "Toolkit works" {
		t.Error("Expected Toolkit to append 'works'")
	}
}
