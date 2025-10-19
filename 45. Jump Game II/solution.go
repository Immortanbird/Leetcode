package main

func jump(nums []int) int {
	dp := make([]int, len(nums))

	for i := 1; i < len(dp); i++ {
		dp[i] = 10000
	}

	for i := range len(nums) {
		for j := i + 1; j < len(nums) && j <= i+nums[i]; j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}

	return dp[len(dp)-1]
}
