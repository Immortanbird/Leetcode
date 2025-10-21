package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := []string{}
	maxn := 0

	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
		maxn = max(maxn, num)
	}

	if maxn == 0 {
		return "0"
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	var sb strings.Builder

	for _, str := range strs {
		sb.WriteString(str)
	}

	return sb.String()
}
