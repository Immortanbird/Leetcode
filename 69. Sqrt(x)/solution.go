package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 1, x/2

	for l <= r {
		mid := l + (r-l)/2
		if x/mid == mid {
			return mid
		} else if x/mid < mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return r
}
