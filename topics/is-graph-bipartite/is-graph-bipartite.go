// Oops, forgot to write comments. Good luck, bro.
//
// @author xuejunc deerhunter0837@gmail.com
// @create 2024/8/2 下午10:58
package is_graph_bipartite

// 广度遍历
func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	bfs := func(root int) bool {
		q := []int{root}
		colors[root] = 1
		for i := 0; i < len(q); i++ {
			u := q[i]
			color := colors[u]
			nextColor := (color)%2 + 1
			for _, v := range graph[u] {
				if colors[v] == color {
					return false
				}
				if colors[v] == 0 {
					colors[v] = nextColor
					q = append(q, v)
				}
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		if colors[i] == 0 && !bfs(i) {
			return false
		}
	}
	return true
}

// 深度遍历
func isBipartite2(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	var dfs func(int, int) bool
	dfs = func(root, color int) bool {
		colors[root] = color
		nextColor := color%2 + 1
		for _, v := range graph[root] {
			if colors[v] == 0 {
				if !dfs(v, nextColor) {
					return false
				}
			} else if colors[v] == color {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		if colors[i] == 0 && !dfs(i) {
			return false
		}
	}
	return true
}
