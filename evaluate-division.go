func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	hashMap := make(map[string]int)
	var numMap [40][40]float64
	var visited [40]bool

	var index int
	res := make([]float64, 0)
	for i, equa := range equations {
		from, to := equa[0], equa[1]
		if _, ok := hashMap[from]; !ok {
			hashMap[from] = index
			index++
		}
		if _, ok := hashMap[to]; !ok {
			hashMap[to] = index
			index++
		}
		numMap[hashMap[to]][hashMap[to]] = 1
		numMap[hashMap[from]][hashMap[from]] = 1

		numMap[hashMap[to]][hashMap[from]] = 1 / values[i]
		numMap[hashMap[from]][hashMap[to]] = values[i]
	}
	var dfs func(from, to, pos int, current float64) bool
	dfs = func(from, to, pos int, current float64) bool {
		if pos == to {
			res = append(res, current*numMap[to][pos])
			return true
		}
		for i := 0; i <= index; i++ {
			if v := numMap[pos][i]; v > 0 && !visited[i] {
				visited[i] = true
				if dfs(from, to, i, (current * v)) {
					visited[i] = false
					return true
				}
				visited[i] = false
			}
		}
		return false
	}
	for _, querie := range queries {
		from, to := querie[0], querie[1]
		if _, ok := hashMap[from]; !ok {
			res = append(res, -1)
			continue
		}
		if _, ok := hashMap[to]; !ok {
			res = append(res, -1)
			continue
		}
		if v := numMap[hashMap[from]][hashMap[to]]; v > 0 {
			res = append(res, v)
			continue
		}
		visited[hashMap[from]] = true
		if !dfs(hashMap[from], hashMap[to], hashMap[from], 1) {
			res = append(res, -1)
		}
		visited[hashMap[from]] = false
	}
	return res
}