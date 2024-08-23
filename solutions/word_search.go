package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	marked := make([][]bool, m)
	for i := range marked {
		marked[i] = make([]bool, n)
	}
	for row := range board {
		for col := range board[row] {
			if word[0] == board[row][col] {
				if find(row, col, 0, board, marked, word) {
					return true
				}
			}
		}
	}
	return false
}

func find(row, col int, idx int, board [][]byte, marked [][]bool, word string) bool {
	var dir = [4][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}
	if idx == len(word)-1 {
		return true
	}
	marked[row][col] = true
	for _, d := range dir {
		r := row + d[0]
		c := col + d[1]
		if inRange(r, len(board)) && inRange(c, len(board[0])) && board[r][c] == word[idx+1] && !marked[r][c] {
			if find(r, c, idx+1, board, marked, word) {
				return true
			}
		}
	}
	marked[row][col] = false
	return false
}

func inRange(a, b int) bool {
	return 0 <= a && a < b
}
