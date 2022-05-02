package intergers

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(1, 3)
	expected := 4

	if ans != expected {
		t.Errorf("Add %d and %d result should be %d", 1, 3, expected)
	}
}
