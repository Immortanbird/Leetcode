package main

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}

	odd := false
	res := 0
	freq := map[rune]int{}
	for _, ch := range s {
		freq[ch]++
	}

	for _, count := range freq {
		if count%2 == 1 {
			odd = true
			res += count - 1
		} else {
			res += count
		}
	}

	if odd {
		res++
	}
	return res
}
