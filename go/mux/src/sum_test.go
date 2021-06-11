package apifarm

import "testing"

func TestSum(t *testing.T) {
	got := Sum(5, 10)
	expected := 15

	if got != expected {
		t.Errorf("Sum failed, expected %v, got %v", expected, got)
	}
}

func TestSquare(t *testing.T) {
	got := Square(5)
	expected := 25

	if got != expected {
		t.Errorf("Square failed, expected %v, got %v", expected, got)
	}
}
