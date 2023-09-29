func canVisitAllRooms(rooms [][]int) bool {
	queue := []int{0}
	visited := map[int]bool{0: true}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			room := queue[0]
			visited[room] = true
			for _, v := range rooms[room] {
				if _, ok := visited[v]; !ok {
					queue = append(queue, v)
				}
			}
			queue = queue[1:]
		}
	}
	return len(visited) == len(rooms)
}