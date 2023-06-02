package day4

import "testing"

func TestFullyContains(t *testing.T) {
	expected := 2
	actual := FullyContains("/test.txt")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}

func TestAllFullyContains(t *testing.T) {
	expected := 4
	actual := AnyContains("/test.txt")

	// Assert
	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
