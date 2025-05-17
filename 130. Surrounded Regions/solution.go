package main

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		if board[i][0] == 'O' {
			board[i][0] = '1'
			check(board, i, 0)
		}
		if board[i][len(board[0])-1] == 'O' {
			board[i][len(board[0])-1] = '1'
			check(board, i, len(board[0])-1)
		}
	}

	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'O' {
			board[0][j] = '1'
			check(board, 0, j)
		}
		if board[len(board)-1][j] == 'O' {
			board[len(board)-1][j] = '1'
			check(board, len(board)-1, j)
		}
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == '1' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func check(board [][]byte, row int, col int) {
	direction := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := [][]int{{row, col}}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]

		for _, d := range direction {
			ii := i + d[0]
			jj := j + d[1]

			if ii >= 0 && ii < len(board) && jj >= 0 && jj < len(board[0]) && board[ii][jj] == 'O' {
				queue = append(queue, []int{ii, jj})
				board[ii][jj] = '1'
			}
		}

		queue = queue[1:]
	}
}
