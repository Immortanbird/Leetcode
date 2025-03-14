package main

func minZeroArray(nums []int, queries [][]int) int {
	differenceArr := make([]int, len(nums)+1)
	prefixSum := 0
	cnt := 0
	i := 0

	for ; i < len(nums); i++ {
		for ; prefixSum+differenceArr[i] < nums[i] && cnt < len(queries); cnt++ {
			if queries[cnt][1]+1 > i {
				differenceArr[max(queries[cnt][0], i)] += queries[cnt][2]
				differenceArr[queries[cnt][1]+1] -= queries[cnt][2]
			}
		}
		prefixSum += differenceArr[i]
		if prefixSum < nums[i] {
			return -1
		}
	}

	return cnt
}
