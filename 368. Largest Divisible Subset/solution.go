package main

import "slices"

func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)

	dp := make([]int, len(nums))
	prev := make([]int, len(nums))
	ll := 0
	lastindex := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		prev[i] = -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > ll {
			ll = dp[i]
			lastindex = i
		}
	}

	res := []int{nums[lastindex]}
	for i := lastindex; prev[i] != -1; i = prev[i] {
		res = append(res, nums[prev[i]])
	}

	return res
}
