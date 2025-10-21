package main

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	res := 0

	for i, j := 0, 0; i < len(s) && j < len(g); i++ {
		if s[i] >= g[j] {
			res++
			j++
		}
	}

	return res
}
