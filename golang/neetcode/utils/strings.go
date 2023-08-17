package utils

import (
	"sort"
)

func GetAlphabeticalString(str string) string {
	runeSlice := []rune(str)

	sort.Slice(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})

	return string(runeSlice)
}
