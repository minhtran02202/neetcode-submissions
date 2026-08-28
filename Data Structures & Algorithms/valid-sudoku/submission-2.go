func isValidSudoku(board [][]byte) bool {
	var rows, cols, squares [9][9]bool
	
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c ++ {
			if board[r][c] == '.' { continue }
			
			val := board[r][c] - '1'

			squaresI := (r/3) * 3 + c/3

			if rows[r][val] || cols[c][val] || squares[squaresI][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			squares[squaresI][val] = true
		}
	}

	return true
}
