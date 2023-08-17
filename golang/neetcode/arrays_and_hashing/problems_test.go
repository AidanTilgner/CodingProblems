package arraysandhashing

import (
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
