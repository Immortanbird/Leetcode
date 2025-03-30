package main

func partitionLabels(s string) []int {
	lastpos := map[byte]int{}
	res := []int{}

	for i := 0; i < len(s); i++ {
		lastpos[s[i]] = max(lastpos[s[i]], i)
	}

	i := 0
	for i < len(s) {
		last := lastpos[s[i]]

		for j := i + 1; j < last; j++ {
			last = max(last, lastpos[s[j]])
		}

		l := last - i + 1
		res = append(res, l)
		i = last + 1
	}

	return res
}
