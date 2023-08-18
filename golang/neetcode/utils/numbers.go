package utils

import "sort"

func GetDescendingNumbers(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j] // Sort in descending order
	})
	return nums
}
