package twopointers

import (
	"reflect"
	"testing"
)

func TestTwoSumTwoInputArraySorted(t *testing.T) {
	case1 := []int{2, 7, 11, 15}
	case1T := 9

	answer1 := TwoSumTwoInputArraySorted(case1, case1T)

	if !reflect.DeepEqual(answer1, []int{1, 2}) {
		t.Errorf("Expected [1, 2], got: %v", answer1)
	}

	case2 := []int{5, 25, 75}
	case2T := 100

	answer2 := TwoSumTwoInputArraySorted(case2, case2T)

	if !reflect.DeepEqual(answer2, []int{2, 3}) {
		t.Errorf("Expected [2, 3], got: %v", answer2)
	}

	case3 := []int{-1, 0}
	case3T := -1

	answer3 := TwoSumTwoInputArraySorted(case3, case3T)

	if !reflect.DeepEqual(answer3, []int{1, 2}) {
		t.Errorf("Expected [1, 2], got: %v", answer3)
	}
}
