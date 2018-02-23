package tests

import "testing"

// Sum sums two integers
func Sum(x int, y int) int {
	return x + y
}

// TestSum test the sum function
func TestSum(t *testing.T) {
	total := Sum(5, 6)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
