package main

func exist(board [][]byte, word string) bool {
	chosen := make([][]bool, len(board))
	for i := range chosen {
		chosen[i] = make([]bool, len(board[0]))
		for j := range chosen[i] {
			chosen[i][j] = false
		}
	}

	for i := range len(board) {
		for j := range len(board[0]) {
			if board[i][j] == word[0] {
				chosen[i][j] = true
				if DFS(board, i, j, 0, word, chosen) {
					return true
				}
				chosen[i][j] = false
			}
		}
	}

	return false
}

func DFS(board [][]byte, i int, j int, current int, word string, chosen [][]bool) bool {
	if current == len(word)-1 {
		return true
	}

	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, d := range directions {
		ii, jj := i+d[0], j+d[1]
		if ii >= 0 && ii < len(board) && jj >= 0 && jj < len(board[0]) && !chosen[ii][jj] && board[ii][jj] == word[current+1] {
			chosen[ii][jj] = true
			if DFS(board, ii, jj, current+1, word, chosen) {
				return true
			}
			chosen[ii][jj] = false
		}
	}

	return false
}
