package hex

import (
	"testing"
)

func TestNewHex(t *testing.T) {
	hex := NewHex(1, -1)
	expected := Hex{1, -1, 0}
	if hex.Q != expected.Q {
		t.Errorf("Expected %v, got %v", expected.Q, hex.Q)
	}
	if hex.R != expected.R {
		t.Errorf("Expected %v, got %v", expected.R, hex.R)
	}
	if hex.S != expected.S {
		t.Errorf("Expected %v, got %v", expected.S, hex.S)
	}
}

func TestGetNeighbour(t *testing.T) {
	hex := MakeHex(1, 1)

	neighbours := []Hex{
		MakeHex(1, 0),
		MakeHex(2, 0),
		MakeHex(2, 1),
		MakeHex(1, 2),
		MakeHex(0, 2),
		MakeHex(0, 1),
	}

	for i, n := range neighbours {
		if n != hex.GetNeighbour(i) {
			t.Errorf("Expected %v, got %v", n, hex.GetNeighbour(i))
		}
	}
}
