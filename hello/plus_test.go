package hello

import "testing"

func TestPlus(t *testing.T) {
	r := Plus(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}
}
