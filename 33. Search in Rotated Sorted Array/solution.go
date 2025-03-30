package main

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	mid := 0

	for i <= j {
		mid = (i + j) / 2
		if mid > 0 && nums[mid] < nums[mid-1] {
			i = mid
			break
		}

		if nums[mid] < nums[0] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	res := binarySearch(nums[:i], target)
	if res == -1 {
		res = binarySearch(nums[i:], target)
		if res != -1 {
			return res + i
		}
	}

	return res
}

func binarySearch(nums []int, target int) int {
	i := 0
	j := len(nums) - 1

	for i <= j {
		mid := (i + j) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return -1
}
