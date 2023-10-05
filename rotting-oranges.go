func orangesRotting(grid [][]int) int {
	var visited [10][10]bool
	queue := make([]int, 0)
	r, c := len(grid)-1, len(grid[0])-1
	var time, orange, appended int
	for i, row := range grid {
		for j, val := range row {
			switch {
			case val == 1:
				orange++
			case val == 2:
				appended++
				queue = append(queue, i, j)
				visited[i][j] = true
			}
		}
	}
	for len(queue) > 0 {
		n := appended
		appended = 0
		if orange == 0 {
			return time
		}
		for i := 0; i < n; i++ {
			row, col := queue[0], queue[1]
			queue = queue[2:]
			col++
			if col <= c && grid[row][col] == 1 && !visited[row][col] {
				visited[row][col] = true
				orange--
				queue = append(queue, row, col)
				appended++
			}
			col -= 2
			if col >= 0 && grid[row][col] == 1 && !visited[row][col] {
				visited[row][col] = true
				orange--
				queue = append(queue, row, col)
				appended++
			}
			col++
			row--
			if row >= 0 && grid[row][col] == 1 && !visited[row][col] {
				visited[row][col] = true
				orange--
				queue = append(queue, row, col)
				appended++
			}
			row += 2
			if row <= r && grid[row][col] == 1 && !visited[row][col] {
				visited[row][col] = true
				orange--
				queue = append(queue, row, col)
				appended++
			}
		}
		time++
	}
	if orange == 0 {
		return time
	}
	return -1
}
