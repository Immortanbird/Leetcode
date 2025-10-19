package main

func canJump(nums []int) bool {
	farthest := 0
	for i := range len(nums) {
		if i > farthest {
			return false
		}
		farthest = max(farthest, nums[i]+i)
	}

	return true
}
