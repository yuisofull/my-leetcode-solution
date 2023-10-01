func minReorder(n int, connections [][]int) int {
	adj := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool)
	}
	var res int
	for _, conn := range connections {
		adj[conn[0]][conn[1]] = true
		adj[conn[1]][conn[0]] = false
	}
	var dfs func(int)
	dfs = func(val int) {
		for i, v := range adj[val] {
			if v == true {
				res++
			}
			delete(adj[val], i)
			delete(adj[i], val)
			dfs(i)
		}
	}
	dfs(0)
	return res
}