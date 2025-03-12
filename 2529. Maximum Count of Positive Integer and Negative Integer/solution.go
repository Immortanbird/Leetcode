package main

func maximumCount(nums []int) int {
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return 0
	}

	if i, ok := lowerBound(nums, 0); ok {
		j, _ := upperBound(nums, 0)
		return max(i, len(nums)-j)
	} else {
		return max(i, len(nums)-i)
	}
}

func lowerBound(nums []int, target int) (int, bool) {
	l, r := 0, len(nums)

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid, true
		} else if nums[mid] < 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l, false
}

func upperBound(nums []int, target int) (int, bool) {
	l, r := -1, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid, true
		} else if nums[mid] < 0 {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l, false
}
