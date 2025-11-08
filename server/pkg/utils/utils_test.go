package utils

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("test", 3)
	expected := []string{"test", "test", "test"}

	if len(repeated) != len(expected) {
		t.Errorf("Expected %v, got %v", len(expected), len(repeated))
	}

	for i, r := range repeated {
		if r != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], r)
		}
	}
}
