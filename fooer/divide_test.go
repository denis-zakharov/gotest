package fooer

import "testing"

func assertEqual(t *testing.T, expected, actual float64) {
	t.Helper() // Mark this function as a helper
	if expected != actual {
		t.Fatalf("expected %v, but got %v", expected, actual)
	}
}

// Test function
func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertEqual(t, 5, result)

	_, err = Divide(10, 0)
	if err == nil {
		t.Fatalf("expected an error, but got none")
	}
}
