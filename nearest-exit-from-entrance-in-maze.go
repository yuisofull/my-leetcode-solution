func nearestExit(maze [][]byte, entrance []int) int {
	var count int
	appended := 1
	r, c := len(maze)-1, len(maze[0])-1
	var visited [100][100]bool
	visited[entrance[0]][entrance[1]] = true
	queue := []int{entrance[0], entrance[1]}
	for len(queue) > 0 {
		n := appended
		appended = 0
		for i := 0; i < n; i++ {
			row, col := queue[0], queue[1]
			queue = queue[2:]
			if row == 0 || col == 0 || row == r || col == c {
				if count > 0 {
					return count
				}
			}
			col++
			if col <= c && maze[row][col] == '.' && !visited[row][col] {
				visited[row][col] = true
				queue = append(queue, row, col)
				appended++
			}
			col -= 2
			if col >= 0 && maze[row][col] == '.' && !visited[row][col] {
				visited[row][col] = true
				queue = append(queue, row, col)
				appended++
			}
			col++
			row--
			if row >= 0 && maze[row][col] == '.' && !visited[row][col] {
				visited[row][col] = true
				queue = append(queue, row, col)
				appended++
			}
			row += 2
			if row <= r && maze[row][col] == '.' && !visited[row][col] {
				visited[row][col] = true
				queue = append(queue, row, col)
				appended++
			}
		}
		count++
	}
	return -1
}