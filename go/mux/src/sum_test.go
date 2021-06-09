package apifarm

import "testing"

func TestSum(t *testing.T) {
	got := sum(5, 10)
	expected := 15

	if got != expected {
		t.Errorf("sum failed, expected %v, got %v", 15, got)
	}
}
