package search

import "testing"

func TestBinarySearch(t *testing.T) {
	case1 := []int{1, 2, 3, 4, 5, 6}
	target1 := 5

	result := BinarySearch(case1, target1)

	if result != 4 {
		t.Errorf("Expected 4, got %v", result)
	}
}
