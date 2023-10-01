func findCircleNum(isConnected [][]int) int {
	var res int
	visited := map[int]bool{}
	var dfs func(int)
	dfs = func(val int) {
		for i, v := range isConnected[val] {
			if v == 1 && visited[i] == false {
				visited[i] = true
				dfs(i)
			}
		}
	}

	for i, _ := range isConnected {
		if visited[i] == false {
			res++
			visited[i] = true
			dfs(i)
		}
	}
	return res
}