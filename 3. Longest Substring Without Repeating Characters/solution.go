package main

func lengthOfLongestSubstring(s string) int {
	index_map := map[byte]int{}

	maxLength := 0
	i, j := 0, 0

	for ; j < len(s); j++ {
		if index, exist := index_map[s[j]]; exist {
			maxLength = max(maxLength, j-i)
			for ; i < index+1; i++ {
				delete(index_map, s[i])
			}
		}
		index_map[s[j]] = j
	}

	maxLength = max(maxLength, j-i)

	return maxLength
}
