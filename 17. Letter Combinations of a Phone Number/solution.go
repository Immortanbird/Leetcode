package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	dial := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	res := []string{}

	backtrack(digits, 0, "", dial, &res)

	return res
}

func backtrack(digits string, i int, temp string, dial map[byte]string, res *[]string) {
	if i == len(digits) {
		*res = append(*res, temp)
		return
	}

	for _, r := range dial[digits[i]] {
		backtrack(digits, i+1, temp+string(r), dial, res)
	}
}
