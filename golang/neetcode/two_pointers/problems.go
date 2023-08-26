package twopointers

import (
	"sort"
)

func TwoSumTwoInputArraySorted(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

func ThreeSum(nums []int) [][]int {
	n := len(nums)

	sort.Ints(nums)

	var result [][]int

	for num1Index := 0; num1Index < n-2; num1Index++ {
		if num1Index > 0 && nums[num1Index] == nums[num1Index-1] {
			continue
		}

		num2Index := num1Index + 1
		num3Index := n - 1

		for num2Index < num3Index {
			sum := nums[num1Index] + nums[num2Index] + nums[num3Index]
			if sum == 0 {
				result = append(result, []int{nums[num1Index], nums[num2Index], nums[num3Index]})

				num3Index--

				for num2Index < num3Index && nums[num3Index] == nums[num3Index+1] {
					num3Index--
				}
			} else if sum > 0 {
				num3Index--
			} else {
				num2Index++
			}
		}
	}

	return result
}
