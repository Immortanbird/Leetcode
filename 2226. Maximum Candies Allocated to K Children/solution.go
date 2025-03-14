package main

func maximumCandies(candies []int, k int64) int {
	l, r := 0, candies[0]
	var sum int64 = 0
	for _, v := range candies {
		sum += int64(v)
		r = max(r, v)
	}

	for l < r {
		mid := l + (r-l)/2

		var shares int64 = 0
		for _, v := range candies {
			shares += int64(v / mid)
		}

		if shares >= k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}
