package utils

import "sort"

func GetDescendingNumbers(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // Sort in descending order
	})
	return nums
}

func IsBetween(n1 int, n2 int, n3 int) bool {
	return n1 <= n3 && n1 >= n2
}
