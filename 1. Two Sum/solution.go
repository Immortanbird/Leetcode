package main

func twoSum(nums []int, target int) []int {
	hashset := map[int]int{}

	for i, num := range nums {
		if e, exist := hashset[target-num]; exist {
			return []int{i, e}
		}
		hashset[num] = i
	}

	return []int{}
}
