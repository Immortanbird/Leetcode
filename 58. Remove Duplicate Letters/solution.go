package main

import (
	"strings"
)

func removeDuplicateLetters(s string) string {
	lastOccurrence := map[byte]int{}

	for i := range len(s) {
		lastOccurrence[s[i]] = i
	}

	stack := []byte{s[0]}
	exist := map[byte]bool{s[0]: true}
	var sb strings.Builder

	for i := 1; i < len(s); i++ {
		if exist[s[i]] {
			continue
		}
		if stack[len(stack)-1] < s[i] {
			stack = append(stack, s[i])
		} else if stack[len(stack)-1] > s[i] {
			for len(stack) > 0 && stack[len(stack)-1] >= s[i] && lastOccurrence[stack[len(stack)-1]] > i {
				exist[stack[len(stack)-1]] = false
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, s[i])
		}
		exist[s[i]] = true
	}

	sb.Write(stack)

	return sb.String()
}
