package search

func BinarySearch(arr []int, target int) int {
	high := len(arr) - 1
	low := 0

	for low <= high {
		mid := (low + high) / 2
		midValue := arr[mid]

		if midValue == target {
			return mid
		} else if midValue < target {
			low = mid + 1
		} else if midValue > target {
			high = mid - 1
		}
	}

	return -1
}
