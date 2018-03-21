package Day_3

func CalculateWierdSums(input int) int {
	var board [11][11]int
	board[5][5] = 1

	x := 5
	y := 5
	dx := 0
	dy := 1

	for i := 0; i < 10*10; i++ {
		if dy < 0 && board[x-1][y] == 0 {
			dx = -1
			dy = 0
		} else if dx < 0 && board[x][y+1] == 0 {
			dx = 0
			dy = 1
		} else if dy > 0 && board[x+1][y] == 0 {
			dx = 1
			dy = 0
		} else if dx > 0 && board[x][y-1] == 0 {
			dx = 0
			dy = -1
		}

		x = x + dx
		y = y + dy

		board[x][y] =
			board[x-1][y] +
				board[x-1][y-1] +
				board[x]  [y-1] +
				board[x+1][y-1] +
				board[x+1][y] +
				board[x+1][y+1] +
				board[x]  [y+1] +
				board[x-1][y+1]

		if board[x][y] > input {
			return board[x][y]
		}
	}

	return -1
}
