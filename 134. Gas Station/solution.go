package main

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	cur := 0
	idx := -1

	for i := range len(gas) {
		sum += gas[i] - cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			idx = -1
			cur = 0
		} else if idx < 0 {
			idx = i
		}
	}

	if sum < 0 {
		return -1
	} else {
		return idx
	}
}
