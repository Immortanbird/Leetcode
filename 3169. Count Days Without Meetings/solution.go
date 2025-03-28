package main

import "slices"

func countDays(days int, meetings [][]int) int {
	slices.SortFunc(meetings, func(a, b []int) int { return a[0] - b[0] })

	idle := 0
	cursor := 1
	for _, meeting := range meetings {
		if cursor < meeting[0] {
			idle += meeting[0] - cursor
		}
		cursor = max(cursor, meeting[1]+1)
	}

	if cursor <= days {
		idle += days + 1 - cursor
	}

	return idle
}
