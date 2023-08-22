package arraysandhashing

import (
	"fmt"
	u "neetcode-solutions/utils"
	"sort"
)

// ValidAnagram takes two strings, and returns true if the second string
// is an anagram of the first string.
// reference: https://leetcode.com/problems/valid-anagram/
func ValidAnagram(s string, t string) bool {
	sletters := make(map[rune]int)
	tletters := make(map[rune]int)

	for _, v := range s {
		sletters[v]++
	}

	for _, v := range t {
		tletters[v]++
	}

	if len(sletters) != len(tletters) {
		return false
	}

	for k, v := range sletters {
		if tletters[k] != v {
			return false
		}
	}

	return true
}

// GroupAnagrams takes a list of strings, and returns a two-dimensional
// list of anagrams derived from the input list.
// reference: https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(words []string) [][]string {
	mappings := make(map[string][]string)

	for _, w := range words {
		al := u.GetAlphabeticalString(w)
		mappings[al] = append(mappings[al], w)
	}

	var results [][]string

	for _, set := range mappings {
		results = append(results, set)
	}

	// time complexity: O(n * m log n)
	// where n is the number of words and m is the max length of a word
	return results
}

func TopKFrequentElements(nums []int, k int) []int {
	mappings := make(map[int]int)

	for _, n := range nums {
		mappings[n]++
	}

	var resultsWithTotals [][2]int

	for v, n := range mappings {
		resultsWithTotals = append(resultsWithTotals, [2]int{n, v})
	}

	sort.Slice(resultsWithTotals, func(i, j int) bool { // n log n worst case
		return resultsWithTotals[i][0] > resultsWithTotals[j][0]
	})

	var results []int

	for i := 0; i < k; i++ {
		results = append(results, resultsWithTotals[i][1])
	}

	return results
}

func ProductOfArrayExceptSelf(nums []int) []int {
	totalReal := 1
	totalWithoutSingleZero := 1
	foundZero := false

	for _, v := range nums {
		totalReal *= v
		if v == 0 && !foundZero {
			foundZero = true
			continue
		}
		totalWithoutSingleZero *= v
	}

	if totalWithoutSingleZero == 0 {
		return make([]int, len(nums))
	}

	if foundZero {
		fmt.Printf("Found zero btw\n")
		results := make([]int, len(nums))

		for i, v := range nums {
			if v == 0 {
				results[i] = totalWithoutSingleZero
				continue
			}
			results[i] = 0
		}

		return results
	}

	results := make([]int, len(nums))
	for i, v := range nums {
		results[i] = totalReal / v
	}

	return results
}

func ValidSodoku([][]byte) bool {
	return false
}
