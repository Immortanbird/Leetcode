package main

import "slices"

func minOperations(nums []int, k int) int {
	slices.Sort(nums)

	if k > nums[0] {
		return -1
	}

	cnt := 0
	if nums[len(nums)-1] > k {
		cnt++
	}
	for i := len(nums) - 2; i >= 0 && nums[i] > k; i-- {
		if nums[i] != nums[i+1] {
			cnt++
		}
	}

	return cnt
}
