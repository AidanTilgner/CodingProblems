package arraysandhashing

import (
	"reflect"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	if !ValidAnagram("anagram", "nagaram") {
		t.Error("Expected true, got false")
	}

	if ValidAnagram("rat", "car") {
		t.Error("Expected false, got true")
	}
}

func TestGroupAnagrams(t *testing.T) {
	val := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	if len(val) != 3 {
		t.Errorf("Expected 3, got %d", len(val))
	}
}

func TestTopKFrequentElements(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	val := TopKFrequentElements(nums, k)

	equal := reflect.DeepEqual(val, []int{1, 2})

	if !equal {
		t.Errorf("Expected [1, 2], got %d", val)
	}
}

func TestProductOfArrayExceptSelf(t *testing.T) {
	cases := [][][]int{
		{{1, 2, 3, 4}, {24, 12, 8, 6}},
		{{-1, 1, 0, -3, 3}, {0, 0, 9, 0, 0}},
	}

	for _, v := range cases {
		val := ProductOfArrayExceptSelf(v[0])

		equal := reflect.DeepEqual(v[1], val)

		if !equal {
			t.Errorf("Expected %v got %v", v[1], val)
		}
	}
}

func TestLongestConsecutiveSequence(t *testing.T) {
	case1 := []int{100, 4, 200, 1, 3, 2}
	res := LongestConsecutiveSequence(case1)
	if res != 4 {
		t.Errorf(
			"Expected length 4 got %v",
			res,
		)
	}
}
