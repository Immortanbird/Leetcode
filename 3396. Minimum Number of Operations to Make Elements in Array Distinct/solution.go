package main

func minimumOperations(nums []int) int {
	hash := map[int]bool{}

	for i := len(nums) - 1; i >= 0; i-- {
		if hash[nums[i]] {
			return i/3 + 1
		}
		hash[nums[i]] = true
	}

	return 0
}
