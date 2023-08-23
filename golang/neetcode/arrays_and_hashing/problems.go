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

func ValidSudoku(board [][]byte) bool { // todo: use fixed arrays instead to reduce space complexity, remove dynamic resizes, etc
	rowSet := make(map[int]map[int]struct{})
	columnSet := make(map[int]map[int]struct{})
	ninthSet := make(map[string]map[int]struct{})

	for r, row := range board {
		for c, rSquare := range row {
			if rSquare == '.' {
				continue
			}
			numSquare := int(rSquare - '0') // Convert to actual integer value
			if numSquare < 1 || numSquare > 9 {
				return false
			}
			if rowSet[r] == nil {
				rowSet[r] = make(map[int]struct{})
			}
			if _, exists := rowSet[r][numSquare]; exists {
				return false
			}
			rowSet[r][numSquare] = struct{}{}
			if columnSet[c] == nil {
				columnSet[c] = make(map[int]struct{})
			}
			if _, exists := columnSet[c][numSquare]; exists {
				return false
			}
			columnSet[c][numSquare] = struct{}{}
			nnIndex := fmt.Sprintf("%d%d", r/3, c/3)
			if ninthSet[nnIndex] == nil {
				ninthSet[nnIndex] = make(map[int]struct{})
			}
			if _, exists := ninthSet[nnIndex][numSquare]; exists {
				return false
			}
			ninthSet[nnIndex][numSquare] = struct{}{}
		}
	}

	return true
}
