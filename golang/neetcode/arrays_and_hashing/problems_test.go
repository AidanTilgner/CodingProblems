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
